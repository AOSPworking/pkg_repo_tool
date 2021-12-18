# pkg_repo_tool

## 1. Usage

把从 `frameworks` 收集到的 `bp` 文件放到一个文件夹中。在这里为了方便叙述，就把这个文件夹命名为 `frameworks` 并且放在 `bprepo2json/main` 之下。

然后：

```shell
> cd bprepo2json/main
> go run .  -l framework.repo.list \
>           -o repo_pkg_module.json \
>           ./frameworks
```

接下来就会在 `bprepo2json/main` 文件夹下输出一个名为 `out.json` 的文件，里面写着具体的关系。

## 2. Misc

* 直接使用了 `Blueprint` 中的 `parser`；
* 我自己多写的那部分代码非常烂。
