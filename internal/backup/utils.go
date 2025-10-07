package backup

import (
	"fmt"
	"os"
	"os/exec"
)

// BackupMySQL runs a mysqldump command for the given database.
func BackupMySQL(user, password, dbName, outputFile string) error {
	fmt.Println("🗄️ Backing up MySQL database:", dbName)
	cmd := exec.Command("mysqldump", "-u", user, fmt.Sprintf("-p%s", password), dbName)
	outFile, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer outFile.Close()
	cmd.Stdout = outFile
	return cmd.Run()
}

// RestoreMySQL restores a MySQL database from a dump file.
func RestoreMySQL(user, password, dbName, inputFile string) error {
	fmt.Println("📦 Restoring MySQL database:", dbName)
	cmd := exec.Command("mysql", "-u", user, fmt.Sprintf("-p%s", password), dbName)
	inFile, err := os.Open(inputFile)
	if err != nil {
		return err
	}
	defer inFile.Close()
	cmd.Stdin = inFile
	return cmd.Run()
}
