# Configurator Plugin Example

This repository is simply demonstrates how to create an external generator plugin for the configurator. The plugin can be modified and built using the following command:

```bash
go build -buildmode=plugin -o testplugin.so testplugin.go
```

To use the plugin with the `configurator`, drop the `testplugin.so` file in a directory under the `plugins` list in the `config.yaml` file. Then, add a target with the template files to use when generating new configs.
