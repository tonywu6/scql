// Copyright 2023 Ant Group Co., Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

#include "engine/datasource/duckdb_wrapper.h"

#include "butil/files/temp_file.h"
#include "fmt/format.h"
#include "gflags/gflags.h"
#include "gtest/gtest.h"

#include "engine/datasource/csvdb_conf.pb.h"

namespace scql::engine {

DECLARE_string(restricted_read_path);
namespace {

::testing::AssertionResult ColumnEquals(
    duckdb::ClientContext& context, const duckdb::DataChunk& chunk,
    size_t col_idx, const std::vector<duckdb::Value>& values) {
  if (chunk.size() != values.size()) {
    return ::testing::AssertionFailure()
           << fmt::format("chunk' size({}) doesn't match value's size({})",
                          chunk.size(), values.size());
  }
  auto& vector = chunk.data[col_idx];
  for (std::size_t i = 0; i < values.size(); i++) {
    if (!duckdb::Value::ValuesAreEqual(context, vector.GetValue(i),
                                       values[i])) {
      return ::testing::AssertionFailure()
             << "the #" << i
             << "th element is not equal, got=" << vector.GetValue(i).ToString()
             << ", while expected=" << values[i].ToString();
    }
  }
  return ::testing::AssertionSuccess();
}

}  // namespace

class DuckdbWrapperTest : public ::testing::Test {
 protected:
  void SetUp() override {
    temp_file_ = std::make_unique<butil::TempFile>("csv");
    temp_file_->save(csv_content_.c_str());
    {
      csvdb_conf_.set_db_name("csvdb");
      auto table = csvdb_conf_.add_tables();
      table->set_table_name("staff");
      table->set_data_path(temp_file_->fname());

      auto column = table->add_columns();
      column->set_column_name("ID");  // for test of ignoring column case
      column->set_column_type(csv::ColumnType::LONG);
      column = table->add_columns();
      column->set_column_name("age");
      column->set_column_type(csv::ColumnType::LONG);
      column = table->add_columns();
      column->set_column_name("name");
      column->set_column_type(csv::ColumnType::STRING);
      column = table->add_columns();
      column->set_column_name("income");
      column->set_column_type(csv::ColumnType::DOUBLE);
    }
  }

 public:
  csv::CsvdbConf csvdb_conf_;
  std::unique_ptr<butil::TempFile> temp_file_;
  std::string csv_content_ = R"csv(ID,age,name,income
1,25,bob,25000.1
2,30,alice,48000
3,42,tom,98000)csv";
};

TEST_F(DuckdbWrapperTest, NormalQuery) {
  FLAGS_restricted_read_path = "./";
  duckdb::DuckDB db = DuckDBWrapper::CreateDB(&csvdb_conf_);
  duckdb::Connection conn(db);

  conn.BeginTransaction();
  DuckDBWrapper::CreateCSVScanFunction(conn);
  conn.Commit();

  auto result = conn.Query("select * from csvdb.staff");
  EXPECT_FALSE(result->HasError());
  ASSERT_EQ(result->RowCount(), 3);
  ASSERT_EQ(result->ColumnCount(), 4);
  auto data_chunk = result->Fetch();

  // column "id"
  EXPECT_TRUE(ColumnEquals(*conn.context, *data_chunk, 0, {1, 2, 3}));
  // column "age"
  EXPECT_TRUE(ColumnEquals(*conn.context, *data_chunk, 1, {25, 30, 42}));
  // column "name"
  EXPECT_TRUE(
      ColumnEquals(*conn.context, *data_chunk, 2, {"bob", "alice", "tom"}));
  // column "income"
  EXPECT_TRUE(
      ColumnEquals(*conn.context, *data_chunk, 3, {25000.1, 48000, 98000}));
}

TEST_F(DuckdbWrapperTest, QueryWithPredicate) {
  duckdb::DuckDB db = DuckDBWrapper::CreateDB(&csvdb_conf_);
  duckdb::Connection conn(db);

  conn.BeginTransaction();
  DuckDBWrapper::CreateCSVScanFunction(conn);
  conn.Commit();

  auto result = conn.Query("select name from csvdb.staff where age > 30");
  EXPECT_FALSE(result->HasError());
  ASSERT_EQ(result->RowCount(), 1);
  ASSERT_EQ(result->ColumnCount(), 1);
  auto data_chunk = result->Fetch();

  // column "name"
  EXPECT_TRUE(ColumnEquals(*conn.context, *data_chunk, 0, {"tom"}));
}

TEST_F(DuckdbWrapperTest, IngoringColumnCase) {
  duckdb::DuckDB db = DuckDBWrapper::CreateDB(&csvdb_conf_);
  duckdb::Connection conn(db);

  conn.BeginTransaction();
  DuckDBWrapper::CreateCSVScanFunction(conn);
  conn.Commit();

  // id -> ID
  auto result = conn.Query("select id from csvdb.staff");
  EXPECT_FALSE(result->HasError());
  ASSERT_EQ(result->RowCount(), 3);
  ASSERT_EQ(result->ColumnCount(), 1);
  auto data_chunk = result->Fetch();

  EXPECT_TRUE(ColumnEquals(*conn.context, *data_chunk, 0, {1, 2, 3}));
}

}  // namespace scql::engine