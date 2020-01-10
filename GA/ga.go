package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

//--- Define Gene -------------------------------------------------------------+
type City struct {
	x float64
	y float64
}

func (city1 City) Distance(city2 City) float64 {
	x := city1.x - city2.x
	y := city1.y - city2.y
	return math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
}

//-----------------------------------------------------------------------------+

//--- Define Quality of Individual --------------------------------------------+
type Route struct {
	route   []City
	fitness float64
}

func (r *Route) Fitness() {
	pathDistance := 0.0
	for i, val := range r.route {
		fromCity := val
		var toCity City
		if i+1 >= len(r.route) {
			toCity = r.route[0]
		} else {
			toCity = r.route[i+1]
		}
		d := fromCity.Distance(toCity)
		pathDistance += d
	}
	r.fitness = 1.0 / pathDistance
}

//-----------------------------------------------------------------------------+

//--- Initialize Population ---------------------------------------------------+
func createRoute(cityList []City) (route []City) {
	//random sample
	n := len(cityList)
	route = make([]City, n)
	indices := rand.Perm(n)
	for i := 0; i < n; i++ {
		route[i] = cityList[indices[i]]
	}

	return
}

func initializePopulation(popSize int, cityList []City) []Route {
	routes := make([][]City, popSize)
	population := make([]Route, popSize)

	for i := 0; i < popSize; i++ {
		routes[i] = createRoute(cityList)
	}

	for i, e := range routes {
		tempRoute := new(Route)
		tempRoute.route = e
		population[i] = *tempRoute
	}
	return population
}

func createCities(numCities int) []City {
	cityList := make([]City, numCities)
	for i := range cityList {
		x := rand.Float64() * 200
		y := rand.Float64() * 200
		cityList[i] = City{x: x, y: y}
	}

	return cityList
}

//-----------------------------------------------------------------------------+

//--- Breeding Poll -----------------------------------------------------------+
func selection(population []Route, elites int) (newGen []int) {
	//evalaute route
	for i := range population {
		population[i].Fitness()
	}

	//comparison sort - greatest to least
	compArr := make([]int, len(population)) //indices
	for i := 0; i < len(compArr)-1; i++ {
		for j := i + 1; j < len(compArr); j++ {
			if population[i].fitness > population[j].fitness {
				compArr[j]++
			} else {
				compArr[i]++
			}
		}
	}

	ranked := make([]int, len(population)) //routes sorted
	for i := range ranked {
		ranked[compArr[i]] = i
	}

	//select new generation
	newGen = make([]int, len(population))
	i := 0 //next gen index
	for i < elites {
		newGen[i] = ranked[i]
		i++
	}

	remaining := rand.Perm(len(population))
	for i < len(population) {
		newGen[i] = remaining[i-elites]
		i++
	}

	return
}

//-----------------------------------------------------------------------------+

//--- Crossover ---------------------------------------------------------------+
func breed(parent1, parent2 Route) Route {
	geneA := int(rand.Float32() * float32(len(parent1.route)))
	geneB := int(rand.Float32() * float32(len(parent1.route)))

	var start, end int
	if geneA > geneB {
		start, end = geneB, geneA
	} else {
		start, end = geneA, geneB
	}
	// fmt.Println(start, end)

	child := new(Route)
	var route []City
	route = append(route, parent1.route[start:end+1]...)
	i := len(route)
	j := 0 //city index to add to route from parent2
	for i < len(parent1.route) {
		if !Contains(route, parent2.route[j]) {
			route = append(route, parent2.route[j])
			i++
		}
		j++
	}
	child.route = route
	if len(child.route) != len(parent1.route) {
		println("Incorrect route length", len(child.route))
	}
	return *child
}

func breedPopulation(newGen []int, elites int, population ...Route) {
	children := make([]Route, len(newGen))

	//keep elites
	for i := 0; i < elites; i++ {
		children[i] = population[newGen[i]]
	}

	pool := rand.Perm(len(newGen))
	// fmt.Println("newgen", len(newGen), len(children))
	child := elites
	for child < len(newGen) {
		parent1 := &population[newGen[pool[child]]]
		parent2 := &population[newGen[pool[len(pool)-child-1]]]
		children[child] = breed(*parent1, *parent2)
		// fmt.Println("Child:", child)
		child++
	}

	//update population in place
	for i := range population {
		population[i] = children[i]
	}
}

func mutate(mxrt float64, population ...Route) {
	for rtIdx := range population {
		for cityIdx := range population[rtIdx].route {
			if rand.Float64() < mxrt {
				route := &population[rtIdx].route
				swapWith := int(rand.Float32() * float32(len(*route)))
				(*route)[cityIdx], (*route)[swapWith] = (*route)[swapWith], (*route)[cityIdx]
			}
		}
	}
}

func main() {
	rand.Seed(0)
	cities := 250
	popSize := 100
	generations := 5000
	elites := int(.45 * float64(popSize))
	mxrt := 1.0 / (float64(cities) * math.Pow(math.Log10(float64(popSize)), 2) / (math.Log10(2)))

	fmt.Println("Cities:", cities, "\npopSize:", popSize, "\nGenerations:", generations)
	fmt.Println("Elites:", elites, "\nMxrt:", mxrt)
	cityList := createCities(cities)
	population := initializePopulation(popSize, cityList)
	results := make([]float64, generations)

	start := time.Now()
	for i := 0; i < generations; i++ {
		newGen := selection(population, elites)
		results[i] = 1 / population[newGen[0]].fitness //add best in generation
		fmt.Println("Generation:", i, results[i])
		breedPopulation(newGen, elites, population...)
		mutate(mxrt, population...)
	}
	fmt.Println("Results:\n", results)
	fmt.Println("Final:", 1/population[selection(population, elites)[0]].fitness)
	fmt.Println("Cities:", cities, "\npopSize:", popSize, "\nGenerations:", generations)
	fmt.Println("Elites:", elites, "\nMxrt:", mxrt)
	elapsed := (time.Now()).Sub(start)
	fmt.Println("Elapsed Time:", elapsed)
	// printPop(false, population...)
}

func printPop(fitness bool, population ...Route) {
	for i := range population {
		if fitness {
			fmt.Println(population[i].fitness)
		} else {
			fmt.Println(population[i])
		}
	}
}

//Contains: returns true is list contains an element == val
//list: list of values to look at
//val:	val to compare elements of list to
func Contains(list []City, val City) bool {
	for _, ele := range list {
		if ele == val {
			return true
		}
	}
	return false
}
