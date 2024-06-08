package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/sajari/regression"
	"gonum.org/v1/gonum/mat"
)

// ReadCSV reads a CSV file and returns the data as a matrix and labels as a slice.
func ReadCSV(filename string) (data *mat.Dense, labels []float64, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, nil, err
	}

	nRows := len(records) - 1
	nCols := len(records[0]) - 1
	rawData := make([]float64, nRows*nCols)
	labels = make([]float64, nRows)

	for i, record := range records[1:] {
		for j, val := range record[:nCols] {
			rawData[i*nCols+j], _ = strconv.ParseFloat(val, 64)
		}
		labels[i], _ = strconv.ParseFloat(record[nCols], 64)
	}

	data = mat.NewDense(nRows, nCols, rawData)
	return data, labels, nil
}

// Shuffle shuffles the data and labels in unison.
func Shuffle(data *mat.Dense, labels []float64) {
	rand.Seed(time.Now().UnixNano())
	nRows, nCols := data.Dims()

	for i := nRows - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		// Swap rows in data
		for k := 0; k < nCols; k++ {
			data.Set(i, k, data.At(j, k))
			data.Set(j, k, data.At(i, k))
		}
		// Swap corresponding labels
		labels[i], labels[j] = labels[j], labels[i]
	}
}

// Split splits the data and labels into training and testing sets.
func Split(data *mat.Dense, labels []float64, testSize float64) (trainData, testData *mat.Dense, trainLabels, testLabels []float64) {
	nRows, nCols := data.Dims()
	nTest := int(float64(nRows) * testSize)
	nTrain := nRows - nTest

	trainData = mat.NewDense(nTrain, nCols, nil)
	testData = mat.NewDense(nTest, nCols, nil)
	trainLabels = make([]float64, nTrain)
	testLabels = make([]float64, nTest)

	for i := 0; i < nTrain; i++ {
		trainData.SetRow(i, data.RawRowView(i))
		trainLabels[i] = labels[i]
	}

	for i := 0; i < nTest; i++ {
		testData.SetRow(i, data.RawRowView(nTrain+i))
		testLabels[i] = labels[nTrain+i]
	}

	return
}

// Model training and evaluation functions would be here...

func main() {
	// Step 2: Read the CSV files
	legitData, legitLabels, err := ReadCSV("structured_data_legitimate.csv")
	if err != nil {
		log.Fatalf("failed to read legitimate data: %v", err)
	}

	phishData, phishLabels, err := ReadCSV("structured_data_phishing.csv")
	if err != nil {
		log.Fatalf("failed to read phishing data: %v", err)
	}

	// Step 3: Combine and shuffle the data
	nLegit, _ := legitData.Dims()
	nPhish, _ := phishData.Dims()

	combinedData := mat.NewDense(nLegit+nPhish, legitData.RawMatrix().Cols, nil)
	combinedData.Stack(legitData, phishData)
	combinedLabels := append(legitLabels, phishLabels...)

	Shuffle(combinedData, combinedLabels)

	// Step 4: Split data into training and testing sets
	trainData, testData, trainLabels, testLabels := Split(combinedData, combinedLabels, 0.2)

	// Step 6: Create and train a model (example with linear regression)
	var r regression.Regression
	r.SetObserved("label")
	r.SetVar(0, "feature1")
	r.SetVar(1, "feature2")
	// Add all features...

	for i := 0; i < trainData.RawMatrix().Rows; i++ {
		features := trainData.RawRowView(i)
		r.Train(regression.DataPoint(trainLabels[i], features))
	}

	r.Run()

	// Step 8: Make predictions and evaluate the model
	predictions := make([]float64, testData.RawMatrix().Rows)
	for i := 0; i < testData.RawMatrix().Rows; i++ {
		features := testData.RawRowView(i)
		pred, _ := r.Predict(features)
		predictions[i] = pred
	}

	// Step 9: Calculate confusion matrix
	tn, fp, fn, tp := confusionMatrix(testLabels, predictions)

	// Step 10: Calculate accuracy, precision, and recall
	accuracy := (tp + tn) / (tp + tn + fp + fn)
	precision := tp / (tp + fp)
	recall := tp / (tp + fn)

	fmt.Printf("accuracy --> %f\n", accuracy)
	fmt.Printf("precision --> %f\n", precision)
	fmt.Printf("recall --> %f\n", recall)
}

// confusionMatrix calculates the confusion matrix.
func confusionMatrix(yTrue, yPred []float64) (tn, fp, fn, tp float64) {
	for i := range yTrue {
		switch {
		case yTrue[i] == 1 && yPred[i] == 1:
			tp++
		case yTrue[i] == 0 && yPred[i] == 1:
			fp++
		case yTrue[i] == 1 && yPred[i] == 0:
			fn++
		case yTrue[i] == 0 && yPred[i] == 0:
			tn++
		}
	}
	return
}

// calculateMeasures calculates accuracy, precision, and recall.
func calculateMeasures(tn, tp, fn, fp float64) (accuracy, precision, recall float64) {
	accuracy = (tp + tn) / (tp + tn + fp + fn)
	precision = tp / (tp + fp)
	recall = tp / (tp + fn)
	return
}
