package prototype

import "testing"

func TestResume(t *testing.T) {
	resume := new(Resume)
	resume.setName("程斌")
	resume.setInfo("男", 21)
	resume.setPhoto("。。。。nan .....")
	resume.setExpericence("2001-2008", "四川阆中XX公司")
	resume.display()

	newResume := resume.clone()
	newResume.setName("朱亚")
	newResume.setPhoto("。。。。。女。。。。")
	newResume.display()

	resume.display()
}
