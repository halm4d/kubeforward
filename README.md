# Kubeforward

## Environments (optional)

### KUBEFORWARD_CONFIG_PATH
Path to config file directory. 

Config file name has to be `config.yaml`.

Default: Executable file directory.

## Example config file
```yaml
kubeConfig: C:\Users\myprofilename\.kube\config # optional - this is also the default value
namespaces:
  - name: mytestingns # namespace
    podRegex: .*nginx.* # regex for pod/service
    ports:
      local: 80
      remote: 80
```