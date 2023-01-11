package main

import "fmt"

func main() {
	fmt.Println("Hello GitHub")
}

//step for use GitHub
/*
	git init => ติดตั้ง git ใน project นั้นๆก่อน ถึงจะใช้คำสั่งอื่นๆได้

	git config => ใช้ในการ set username ที่จะแสดงบน origin
	ex. git config --global user.name "Flxke"
		git config --global user.email "Flxke@gmail.com" email ที่ใช้คือที่สมัครกับ git hub
		git config --global user.password "88888888" กรณีืที่ใช้ git เป็น https

	git clone => ใช้ในกรณีที่ไม่มี repo นั้นมาก่อน
	ex. git clone url

	git status => ตรวจสอบสถานะของแฟ้ม ว่ามีการเปลี่ยนแปลงบ้าง

	git add => เพิ่มไฟล์ที่มีการเปลี่ยนแปลงก่อนเพื่อจะ commit
	ex. git add .
		git add url

	git commit => บอกว่าทำอะไรไปบ้าง
	ex. git commit -m "message"

	git push => ส่งข้อมูลที่ commit ขึ้นไปบน repo
	ex. git push origin [branch name]

	git pull => การดึงข้อมูล remote repo file มายัง local repo

	git switch [branch name]

	git checkout [path or branch name] => คือการที่ต้องออกจะออกจากอะไรสักอย่าง เปลี่ยน branch
	ex. git checkout app.js
		git checkout develop

	git merge [branch name] => รวม เช่น เราทำงานที่ branch develop ต้องการเอาไปรวมกับ master ในการ merge
	ไปที่ branch master แล้วใช้คำสั่ง git merge develop

	git fetch => ใช้ดึงการเปลี่ยนแปลงจาก remote repo มายัง local repo

	git stash => การ save การเปลี่ยนแปลงของเราไว้ก่อน สามารถ restore กลับได้ทุกเมื่อ
	ex. git stash save "comment message"
		git stash pop [index]
		git stash list
		git stash drop [index]

*/
