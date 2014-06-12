package merge_strings_test

import (
	"os"
	"path/filepath"

	. "github.com/maximilien/i18n4cf/integration/test_helpers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("merge-strings -d dirName -r", func() {
	var (
		outputDir         string
		rootPath          string
		fixturesPath      string
		inputFilesPath    string
		expectedFilesPath string
	)

	BeforeEach(func() {
		dir, err := os.Getwd()
		Ω(err).ShouldNot(HaveOccurred())
		rootPath = filepath.Join(dir, "..", "..")
		outputDir = filepath.Join(rootPath, "tmp")

		fixturesPath = filepath.Join("..", "..", "test_fixtures", "merge_strings")
		inputFilesPath = filepath.Join(fixturesPath, "r_option", "input_files")
		expectedFilesPath = filepath.Join(fixturesPath, "r_option", "expected_output")
	})

	Context("can combine multiple language files per directory", func() {
		BeforeEach(func() {
			session := Runi18n("-c", "merge-strings", "-v", "-r", "-d", filepath.Join(inputFilesPath), "--source-language", "en")
			Ω(session.ExitCode()).Should(Equal(0))
		})

		AfterEach(func() {
			RemoveAllFiles(
				GetFilePath(inputFilesPath, "en.all.json"),
				GetFilePath(inputFilesPath+"/sub", "en.all.json"),
			)
		})

		It("en.all.json contains translations from both files", func() {
			CompareExpectedToGeneratedTraslationJson(
				GetFilePath(expectedFilesPath, "en.all.json"),
				GetFilePath(inputFilesPath, "en.all.json"),
			)
			CompareExpectedToGeneratedTraslationJson(
				GetFilePath(expectedFilesPath+"/sub", "en.all.json"),
				GetFilePath(inputFilesPath+"/sub", "en.all.json"),
			)
		})
	})

})
