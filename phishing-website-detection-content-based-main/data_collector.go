package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

type DataRow struct {
	HasTitle             bool
	HasInput             bool
	HasButton            bool
	HasImage             bool
	HasSubmit            bool
	HasLink              bool
	HasPassword          bool
	HasEmailInput        bool
	HasHiddenElement     bool
	HasAudio             bool
	HasVideo             bool
	NumberOfInputs       int
	NumberOfButtons      int
	NumberOfImages       int
	NumberOfOption       int
	NumberOfList         int
	NumberOfTh           int
	NumberOfTr           int
	NumberOfHref         int
	NumberOfParagraph    int
	NumberOfScript       int
	LengthOfTitle        int
	HasH1                bool
	HasH2                bool
	HasH3                bool
	LengthOfText         int
	NumberOfClickableBtn int
	NumberOfA            int
	NumberOfImg          int
	NumberOfDiv          int
	NumberOfFigure       int
	HasFooter            bool
	HasForm              bool
	HasTextArea          bool
	HasIframe            bool
	HasTextInput         bool
	NumberOfMeta         int
	HasNav               bool
	HasObject            bool
	HasPicture           bool
	NumberOfSources      int
	NumberOfSpan         int
	NumberOfTable        int
	URL                  string
	Label                int
}

func mai() { //ubah jadi main() untuk jalanin
	urlFileName := "verified_online_2.csv"
	urlList, err := readCSV(urlFileName)
	if err != nil {
		log.Fatal(err)
	}

	begin, end := 35000, 40000
	collectionList := urlList[begin:end]

	data := createStructuredData(collectionList)

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
		"url",
		"label",
	}

	writeCSV("structured_data_phishing_2.csv", columns, data)
}

// baca url dari csv dan return string
func readCSV(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	urls := make([]string, len(records)-1)
	for i, record := range records[1:] {
		urls[i] = record[0]
	}
	return urls, nil
}

// ambil data dan buat jadi terstruktur
func createStructuredData(urlList []string) []DataRow {
	dataList := []DataRow{}
	for i, url := range urlList {
		fmt.Println("Processing URL:", i, url)
		resp, err := http.Get(url)
		if err != nil || resp.StatusCode != http.StatusOK {
			fmt.Printf("%d. HTTP connection was not successful for the URL: %s\n", i, url)
			continue
		}
		defer resp.Body.Close()

		buf := new(bytes.Buffer)
		buf.ReadFrom(resp.Body)
		htmlContent := buf.String()

		features, err := getFeaturesFromServer(htmlContent)
		if err != nil {
			fmt.Printf("%d. Error extracting features for the URL: %s\n", i, url)
			continue
		}

		vector := createVector(features)
		vector.URL = url
		vector.Label = 1 // Assuming all URLs are phishing for this example

		dataList = append(dataList, vector)
	}
	return dataList
}

// send html to feature_extraction_server dan ambil response
func getFeaturesFromServer(htmlContent string) ([]int, error) {
	resp, err := http.PostForm("http://localhost:8080/extract_features", url.Values{"html": {htmlContent}})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var featureResponse struct {
		Features []int `json:"features"`
	}
	err = json.NewDecoder(resp.Body).Decode(&featureResponse)
	if err != nil {
		return nil, err
	}

	return featureResponse.Features, nil
}

// map features ke datarow
func createVector(features []int) DataRow {
	return DataRow{
		HasTitle:             features[0] == 1,
		HasInput:             features[1] == 1,
		HasButton:            features[2] == 1,
		HasImage:             features[3] == 1,
		HasSubmit:            features[4] == 1,
		HasLink:              features[5] == 1,
		HasPassword:          features[6] == 1,
		HasEmailInput:        features[7] == 1,
		HasHiddenElement:     features[8] == 1,
		HasAudio:             features[9] == 1,
		HasVideo:             features[10] == 1,
		NumberOfInputs:       features[11],
		NumberOfButtons:      features[12],
		NumberOfImages:       features[13],
		NumberOfOption:       features[14],
		NumberOfList:         features[15],
		NumberOfTh:           features[16],
		NumberOfTr:           features[17],
		NumberOfHref:         features[18],
		NumberOfParagraph:    features[19],
		NumberOfScript:       features[20],
		LengthOfTitle:        features[21],
		HasH1:                features[22] == 1,
		HasH2:                features[23] == 1,
		HasH3:                features[24] == 1,
		LengthOfText:         features[25],
		NumberOfClickableBtn: features[26],
		NumberOfA:            features[27],
		NumberOfImg:          features[28],
		NumberOfDiv:          features[29],
		NumberOfFigure:       features[30],
		HasFooter:            features[31] == 1,
		HasForm:              features[32] == 1,
		HasTextArea:          features[33] == 1,
		HasIframe:            features[34] == 1,
		HasTextInput:         features[35] == 1,
		NumberOfMeta:         features[36],
		HasNav:               features[37] == 1,
		HasObject:            features[38] == 1,
		HasPicture:           features[39] == 1,
		NumberOfSources:      features[40],
		NumberOfSpan:         features[41],
		NumberOfTable:        features[42],
	}
}

// bikin csv dari data
func writeCSV(filename string, columns []string, data []DataRow) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if fileInfo, err := file.Stat(); err == nil && fileInfo.Size() == 0 {
		if err := writer.Write(columns); err != nil {
			log.Fatal(err)
		}
	}

	for _, row := range data {
		record := []string{
			strconv.FormatBool(row.HasTitle),
			strconv.FormatBool(row.HasInput),
			strconv.FormatBool(row.HasButton),
			strconv.FormatBool(row.HasImage),
			strconv.FormatBool(row.HasSubmit),
			strconv.FormatBool(row.HasLink),
			strconv.FormatBool(row.HasPassword),
			strconv.FormatBool(row.HasEmailInput),
			strconv.FormatBool(row.HasHiddenElement),
			strconv.FormatBool(row.HasAudio),
			strconv.FormatBool(row.HasVideo),
			strconv.Itoa(row.NumberOfInputs),
			strconv.Itoa(row.NumberOfButtons),
			strconv.Itoa(row.NumberOfImages),
			strconv.Itoa(row.NumberOfOption),
			strconv.Itoa(row.NumberOfList),
			strconv.Itoa(row.NumberOfTh),
			strconv.Itoa(row.NumberOfTr),
			strconv.Itoa(row.NumberOfHref),
			strconv.Itoa(row.NumberOfParagraph),
			strconv.Itoa(row.NumberOfScript),
			strconv.Itoa(row.LengthOfTitle),
			strconv.FormatBool(row.HasH1),
			strconv.FormatBool(row.HasH2),
			strconv.FormatBool(row.HasH3),
			strconv.Itoa(row.LengthOfText),
			strconv.Itoa(row.NumberOfClickableBtn),
			strconv.Itoa(row.NumberOfA),
			strconv.Itoa(row.NumberOfImg),
			strconv.Itoa(row.NumberOfDiv),
			strconv.Itoa(row.NumberOfFigure),
			strconv.FormatBool(row.HasFooter),
			strconv.FormatBool(row.HasForm),
			strconv.FormatBool(row.HasTextArea),
			strconv.FormatBool(row.HasIframe),
			strconv.FormatBool(row.HasTextInput),
			strconv.Itoa(row.NumberOfMeta),
			strconv.FormatBool(row.HasNav),
			strconv.FormatBool(row.HasObject),
			strconv.FormatBool(row.HasPicture),
			strconv.Itoa(row.NumberOfSources),
			strconv.Itoa(row.NumberOfSpan),
			strconv.Itoa(row.NumberOfTable),
			row.URL,
			strconv.Itoa(row.Label),
		}
		if err := writer.Write(record); err != nil {
			log.Fatal(err)
		}
	}
}
