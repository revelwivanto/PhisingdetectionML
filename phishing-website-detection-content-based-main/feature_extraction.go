package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"features/features" // Replace with the correct import path for your features package

	"github.com/PuerkitoBio/goquery"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

func main() {
	// Define folder name
	folderName := "mini_dataset"

	// Create the dataframe
	df, err := createDataFrame(folderName)
	if err != nil {
		log.Fatalf("Failed to create DataFrame: %v", err)
	}

	// Print the dataframe
	fmt.Println(df)
}

func openFile(fileName string) (string, error) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		return "", fmt.Errorf("failed to read file %s: %w", fileName, err)
	}
	return string(content), nil
}

func createDoc(text string) (*goquery.Document, error) {
	return goquery.NewDocumentFromReader(strings.NewReader(text))
}

func createVector(doc *goquery.Document) ([]int, error) {
	// Call feature extraction functions from your custom package
	return []int{
		features.HasTitle(doc),
		features.HasInput(doc),
		features.HasButton(doc),
		features.HasImage(doc),
		features.HasSubmit(doc),
		features.HasLink(doc),
		features.HasPassword(doc),
		features.HasEmailInput(doc),
		features.HasHiddenElement(doc),
		features.HasAudio(doc),
		features.HasVideo(doc),
		features.NumberOfInputs(doc),
		features.NumberOfButtons(doc),
		features.NumberOfImages(doc),
		features.NumberOfOption(doc),
		features.NumberOfList(doc),
		features.NumberOfTH(doc),
		features.NumberOfTR(doc),
		features.NumberOfHref(doc),
		features.NumberOfParagraph(doc),
		features.NumberOfScript(doc),
		features.LengthOfTitle(doc),
		features.HasH1(doc),
		features.HasH2(doc),
		features.HasH3(doc),
		features.LengthOfText(doc),
		features.NumberOfClickableButton(doc),
		features.NumberOfA(doc),
		features.NumberOfImg(doc),
		features.NumberOfDiv(doc),
		features.NumberOfFigure(doc),
		features.HasFooter(doc),
		features.HasForm(doc),
		features.HasTextArea(doc),
		features.HasIFrame(doc),
		features.HasTextInput(doc),
		features.NumberOfMeta(doc),
		features.HasNav(doc),
		features.HasObject(doc),
		features.HasPicture(doc),
		features.NumberOfSources(doc),
		features.NumberOfSpan(doc),
		features.NumberOfTable(doc),
	}, nil
}

func createDataFrame(folderName string) (dataframe.DataFrame, error) {
	columns := []string{
		"has_title",
		"has_input",
		"has_button",
		"has_image",
		"has_submit",
		"has_link",
		"has_password",
		"has_email_input",
		"has_hidden_element",
		"has_audio",
		"has_video",
		"number_of_inputs",
		"number_of_buttons",
		"number_of_images",
		"number_of_option",
		"number_of_list",
		"number_of_th",
		"number_of_tr",
		"number_of_href",
		"number_of_paragraph",
		"number_of_script",
		"length_of_title",
		"has_h1",
		"has_h2",
		"has_h3",
		"length_of_text",
		"number_of_clickable_button",
		"number_of_a",
		"number_of_img",
		"number_of_div",
		"number_of_figure",
		"has_footer",
		"has_form",
		"has_text_area",
		"has_iframe",
		"has_text_input",
		"number_of_meta",
		"has_nav",
		"has_object",
		"has_picture",
		"number_of_sources",
		"number_of_span",
		"number_of_table",
	}

	var data [][]int
	folderPath := filepath.Join(".", folderName)
	files, err := os.ReadDir(folderPath)
	if err != nil {
		return dataframe.DataFrame{}, fmt.Errorf("failed to read directory %s: %w", folderName, err)
	}
	for _, file := range files {
		if file.IsDir() || filepath.Ext(file.Name()) != ".html" {
			continue
		}
		filePath := filepath.Join(folderPath, file.Name())
		content, err := openFile(filePath)
		if err != nil {
			log.Printf("Skipping file %s: %v\n", filePath, err)
			continue
		}
		doc, err := createDoc(content)
		if err != nil {
			log.Printf("Skipping file %s: %v\n", filePath, err)
			continue
		}
		vector, err := createVector(doc)
		if err != nil {
			log.Printf("Skipping file %s: %v\n", filePath, err)
			continue
		}
		data = append(data, vector)
	}

	// Prepare series
	seriesList := make([]series.Series, len(columns))
	for i, col := range columns {
		columnData := make([]interface{}, len(data))
		for j, row := range data {
			columnData[j] = row[i]
		}
		seriesList[i] = series.New(columnData, series.Int, col)
	}

	// Create DataFrame
	df := dataframe.New(seriesList...)
	return df, nil
}
