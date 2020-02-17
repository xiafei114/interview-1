1. `git config --global core.editor vim #修改git使用的编辑器`, 模式使用`VISUAL` 或者 `EDIROT` 的环境变量, 若没有设置, 默认使用`vi` 
2. `reset` 撤销本地分支提交, 但无法影响到远程分支, `revert` 新建一个分支, 并推送到远程
3. `git config --get push.default` 查看push默认的远程地址
4. 设置远程分支追踪的两种方式
	- `git checkout -b <branch> <remote addr>/<branch>` 
	- `git branch -u <remote addr>/<branch> <local branch>`
5. `git push <remote> <source>:<destination>` source = 本地分支
6. `git fetch <remote> <destination>:<source>` source  = 本地分支 
7. `git push <remote> :<target>` 删除远程的<target>分支
8. `git fetch <remote> :<target>` 在本地创建<target>分支



PS: 
> 
	https://learngitbranching.js.org/  git学习网站, 主要涉及rebase, cherry-pick, fetch, push, pull, merge, revert
