package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"math/rand"
	"reflect"
	"sort"
	"strings"
	"time"
)

type population [][]string
type popFit struct {
	chr []string
	fit float64
}

var (
	alphabet = strings.Split("abcdefghijklmnopqrstuvwxyz", "")
	target   = strings.Split("easy", "")
	// we keep the 30% of best fitting
	gradedRetainPercent = 0.4
	// we keep randomly 20% of reminded elements
	nonGradedRetainPercent = 0.4
	// size of the population
	populationSize = 100
	// mutation probability
	mutationProb = 0.05
)

func getShaFitness(chrom []string) float64 {
	var fitness int
	str := strings.Join(chrom[:], "")

	h := sha1.New()
	io.WriteString(h, str)
	digest := h.Sum(nil)
	hexdigestList := fmt.Sprintf("% x", digest)
	hexdigest := strings.Join(strings.Split(hexdigestList, " ")[:], "")
	hex := strings.Split(hexdigest, "")

	for i := 0; i < len(target); i++ {
		if target[i] == hex[i] {
			fitness++
		}
	}
	return float64(fitness) / float64(len(target))
}

func getFitness(chrom []string) float64 {
	var fitness int
	for i := 0; i < len(target); i++ {
		if chrom[i] == target[i] {
			fitness++
		}
	}
	return float64(fitness) / float64(len(target))
}

func newChromosome(size int) []string {
	now := time.Now()
	rand.Seed(now.UnixNano())
	chromo := []string{}
	for i := 0; i < size; i++ {
		chromo = append(chromo, alphabet[rand.Intn(len(alphabet))])
	}
	return chromo
}

func newPopulation(popSize, chrSize int) population {
	var pop population
	for i := 1; i < popSize; i++ {
		pop = append(pop, newChromosome(chrSize))
	}
	return pop
}

func selection(pop population) population {
	now := time.Now()
	rand.Seed(now.UnixNano())
	var populationWithFitness []popFit
	var selectedPopulation population
	var fit float64

	for _, chr := range pop {
		fit = getFitness(chr)
		populationWithFitness = append(populationWithFitness, popFit{chr: chr, fit: fit})
	}
	sort.Slice(populationWithFitness, func(i, j int) bool {
		return populationWithFitness[i].fit > populationWithFitness[j].fit
	})
	var element []string

	numberOfElementToKeep := int(gradedRetainPercent*float64(len(pop))) + 1

	for i := 0; i < numberOfElementToKeep; i++ {
		selectedPopulation = append(selectedPopulation, populationWithFitness[i].chr)
	}

	numberOfRandomElement := int(nonGradedRetainPercent*float64(len(pop))) + 1
	for i := 0; i < numberOfRandomElement; i++ {
		element = populationWithFitness[int(rand.Intn(len(pop)))].chr
		selectedPopulation = append(selectedPopulation, element)
	}
	shuffle(selectedPopulation)
	return selectedPopulation
}

func contains(pop population, chr []string) bool {
	for _, p := range pop {
		if reflect.DeepEqual(p, chr) {
			return true
		}
	}
	return false
}

func crossover(chr1, chr2 []string) []string {
	now := time.Now()
	rand.Seed(now.UnixNano())
	var child []string
	for i := 0; i < len(chr1); i++ {
		if rand.Float64() < 0.5 {
			child = append(child, chr1[i])
		} else {
			child = append(child, chr2[i])
		}
	}
	if len(child) != len(target) {
		log.Fatal("unable to crossover: ", child, target)
	}
	return child
}

func mutation(chr []string) []string {
	now := time.Now()
	rand.Seed(now.UnixNano())

	chr[rand.Intn(len(chr))] = alphabet[rand.Intn(len(alphabet))]

	return chr
}

func isResponse(pop population, target []string) bool {
	for _, p := range pop {
		if reflect.DeepEqual(target, p) {
			return true
		}
	}
	return false
}

func nextPopulation(pop population) population {
	selectedPopulation := selection(pop)

	for len(selectedPopulation) < populationSize {
		now := time.Now()
		rand.Seed(now.UnixNano())
		selectedPopulation = append(selectedPopulation, crossover(selectedPopulation[rand.Intn(len(selectedPopulation))], selectedPopulation[rand.Intn(len(selectedPopulation))]))
		if rand.Float64() < mutationProb {
			mutationIndex := rand.Intn(len(selectedPopulation))
			mutation(selectedPopulation[mutationIndex])
		}
	}
	return selectedPopulation
}

func inArray(a []string, s string) bool {
	for _, i := range a {
		if i == s {
			return true
		}
	}
	return false
}

func shuffle(slice population) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for len(slice) > 0 {
		n := len(slice)
		randIndex := r.Intn(n)
		slice[n-1], slice[randIndex] = slice[randIndex], slice[n-1]
		slice = slice[:n-1]
	}
}

func main() {
	pop := newPopulation(populationSize, len(target))
	log.Println(pop)
	var epoch int
	for !isResponse(pop, target) {
		pop = nextPopulation(pop)
		epoch++
	}
	log.Println(pop, epoch)
}
