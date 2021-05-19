package commands

func HelpGetInfo() string {
	content := "scmd git co master 	=>	git checkout master => checkout master branch\n"
	content += "scmd git dif		=>  git diff 			=> checks difference between working area and index\n"
	content += "scmd git diff		=>  git diff --cached 	=> checks difference between index and repository\n"
	content += "scmd git a			=>  git add .		 	=> adds everything to index\n"
	content += "scmd git b			=>  git branch		 	=> show branches\n"

	return content
}
