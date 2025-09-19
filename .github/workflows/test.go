import great_expectations as ge
from ruamel.yaml import YAML

# Assuming you have a DataContext already set up
context = go.get_context()

# Example YAML configuration for a Datasource
yaml_config = """
name: my_sqlite_datasource
class_name: Datasource
module_name: great_expectations.datasource
execution_engine:
  class_name: SqlAlchemyExecutionEngine
  connection_string: sqlite:///data/my_sqlite_database.db
data_connectors:
  default_runtime_data_connector_name:
    class_name: RuntimeDataConnector
    batch_identifiers:
      - default_identifier_name
"""

# Test the YAML configuration
context.test_yaml_config(yaml_config=yaml_config)
