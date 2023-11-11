package main

import (
  "fmt"
  "sync"

)

// Interface for Ice Cream
type IceCream interface {
  getDescription() string
  getCost() int
}

type BasicIceCream struct {
  description string
  cost        int
}

func (bic *BasicIceCream) getDescription() string {
  return bic.description
}

func (bic *BasicIceCream) getCost() int {
  return bic.cost
}

// Base ice cream types
type CreamyIceCream struct {
  BasicIceCream
}

func newCreamyIceCream() IceCream {
  return &CreamyIceCream{
    BasicIceCream: BasicIceCream{
      description: "Creamy Ice Cream",
      cost:        200,
    },
  }
}

type ChocolateIceCream struct {
  BasicIceCream
}

func newChocolateIceCream() IceCream {
  return &ChocolateIceCream{
    BasicIceCream: BasicIceCream{
      description: "Chocolate Ice Cream",
      cost:        200,
    },
  }
}

type StrawberryIceCream struct {
  BasicIceCream
}

func newStrawberryIceCream() IceCream {
  return &StrawberryIceCream{
    BasicIceCream: BasicIceCream{
      description: "Strawberry Ice Cream",
      cost:        200,
    },
  }
}

//Factory
func getIceCream(iceCreamType int) (IceCream, error) {
  if iceCreamType == 1 {
    return newChocolateIceCream(), nil
  } else if iceCreamType == 2 {
    return newCreamyIceCream(), nil
  } else if iceCreamType == 3 {
    return newStrawberryIceCream(), nil
  }
  return nil, fmt.Errorf("Invalid Ice cream type: Cannot create ice cream in the factory")
}


// Additional decorators
type NutsDecorator struct {
  iceCream IceCream
}

func (d *NutsDecorator) getDescription() string {
  return "\nwith Nuts"
}

func (d *NutsDecorator) getCost() int {
  cost := d.iceCream.getCost() + 50
  return cost
}

type ChocolateChipsDecorator struct {
  iceCream IceCream
}

func (d *ChocolateChipsDecorator) getDescription() string {

  return "\nwith Chocolate Chips"
}

func (d *ChocolateChipsDecorator) getCost() int {
  cost := d.iceCream.getCost() + 75
  return cost
}

type FruitDecorator struct {
  iceCream IceCream
}

func (d *FruitDecorator) getDescription() string {
  return "\nwith Fruit"
}

func (d *FruitDecorator) getCost() int {
  cost := d.iceCream.getCost() + 60
  return cost
}

type MaraschinoDecorator struct {
  iceCream IceCream
}

func (d *MaraschinoDecorator) getDescription() string {

  return "\nwith Maraschino Cherry"
}

func (d *MaraschinoDecorator) getCost() int {
  cost := d.iceCream.getCost() + 30
  return cost
}


//Facade
type IceCreamFacade struct {
	userMoney   int
	totalOrders int
}

func NewIceCreamFacade(userMoney int) *IceCreamFacade {
	return &IceCreamFacade{
		userMoney: userMoney,
	}
}

func (f *IceCreamFacade) OrderIceCream(flavor IceCream, addToppings []IceCream) {
	totalCost := calculateTotalCost(flavor, addToppings)
	discount := getInstance(totalCost).total

	if f.userMoney < (totalCost - discount) {
    fmt.Println("-------------------------------------------")
		fmt.Println("Not enough money to buy another ice cream")
    fmt.Println("-------------------------------------------")
	} else {
		f.totalOrders++
		printOrderDetails(flavor, addToppings, totalCost, discount, f.totalOrders)
		f.userMoney -= totalCost - discount
	}
}
func calculateTotalCost(flavor IceCream, addToppings []IceCream) int {
	totalCost := flavor.getCost()
	for _, topping := range addToppings {
		totalCost += topping.getCost()
	}
	totalCost -= flavor.getCost() * len(addToppings)
	return totalCost
}

func hasEnoughMoney(userMoney, totalCost int) bool {
	return userMoney >= totalCost
}

func printOrderDetails(flavor IceCream, addToppings []IceCream, totalCost, discount, totalOrders int) {
	descript := flavor.getDescription()
	description := ""

	for _, topping := range addToppings {
		description += topping.getDescription()
	}

	fmt.Printf("You bought a %s for $%d\n", descript+description, totalCost-discount)
	fmt.Println("Total orders:", totalOrders)

}

func (f *IceCreamFacade) GetRemainingMoney() int {
	return f.userMoney
}


//Singleton
type FirstOrderDiscount struct {
	total int 
}

var discountInstance *FirstOrderDiscount
var discountLock sync.Mutex

func getInstance(totalCost int) *FirstOrderDiscount {
	discountLock.Lock()
	defer discountLock.Unlock()

	if discountInstance == nil {
		discount := int(0.3 * float64(totalCost))
    fmt.Println("-----------------------------------")
		fmt.Println(  "For the first order  discount 30%")
		fmt.Println("-----------------------------------")
		discountInstance = &FirstOrderDiscount{total: discount}
	} else {
    discount := 0
    discountInstance = &FirstOrderDiscount{total: discount}
		
	}

	return discountInstance
}

func main() {
  fmt.Println("-----------------------------------")
  fmt.Println("  Welcome to the ice cream store!")
  fmt.Println("-----------------------------------")
  var userMoney int

  for { 
    fmt.Print("Enter the amount of money: ") 
    fmt.Scanln(&userMoney) 
    
    if userMoney >= 200 { 
     break 
    } else { 
     fmt.Println("You don't have enough money to buy any ice cream. Should be more than 200") 
    } 
     }

  iceCreamStore:= NewIceCreamFacade(userMoney)

  for {
    var choice int

    fmt.Println("1 ------------ Assemble ice cream")
    fmt.Println("2 ------------ Classic ice creams")
    fmt.Println("3 ------------ Exit")
    fmt.Print("Select an option: ")
    fmt.Scanln(&choice)

    if choice == 2 {
      for {
        // Offer classical ice cream options
        fmt.Println("Classic Ice Creams:")
        fmt.Println("1 ------------ Chocolate Ice Cream with Nuts ($250)")
        fmt.Println("2 ------------ Creamy Ice Cream with Chocolate Chips ($275)")
        fmt.Println("3 ------------ Strawberry Ice Cream with Fruit ($260)")
        fmt.Println("4 ------------ Back")
        fmt.Print("Select a classical ice cream flavor: ")
        fmt.Scanln(&choice)

        var flavor IceCream
        var toppings []IceCream
        backToMainMenu := false

        switch choice {
        case 1:
          flavor, _ = getIceCream(1)
          toppings = append(toppings, &NutsDecorator{flavor})
        case 2:
          flavor, _ = getIceCream(2)
          toppings = append(toppings, &ChocolateChipsDecorator{flavor})
        case 3:
          flavor, _ = getIceCream(3)
          toppings = append(toppings, &FruitDecorator{flavor})
        case 4:
          backToMainMenu = true 
        break
        default:
          fmt.Println("Invalid choice")
          continue
        }
     
      if backToMainMenu {
        fmt.Println("Going back to main menu.")
        break
      }

 
          iceCreamStore.OrderIceCream(flavor, toppings)
          iceCreamRemaining := iceCreamStore.GetRemainingMoney()
          fmt.Println("Your Remaining Money: ", iceCreamRemaining)
          fmt.Println("Do you want to buy another classic ice cream? (1 - Yes, 2 - No): ")
        
          fmt.Scanln(&choice)
  
          if choice == 2 {
            
            break
  
        }
   
      }
      continue
    }

    if choice == 3 {
      fmt.Println("Have a nice day!")
      break
    }

    if choice != 1 && choice != 2{
      fmt.Println("Wrong choice")
      continue
    }

    fmt.Println("1 ------------ Chocolate ice cream ($200)")
    fmt.Println("2 ------------ Creamy ice cream ($200)")
    fmt.Println("3 ------------ Strawberry ice cream ($200)")
    fmt.Println("4 ------------ Back")
    fmt.Print("Select your ice cream flavor: ")
    fmt.Scanln(&choice)

    var flavor IceCream

    switch choice {
    case 1:
   
      flavor, _ = getIceCream(1)
    case 2:
      flavor, _ = getIceCream(2)
    case 3:
      flavor, _ = getIceCream(3)
    case 4:
      continue
    default:
      fmt.Println("Invalid choice")
      continue
    }

    var toppings []IceCream

    for {

      fmt.Println("1 ------------ Nuts ($50)")
      fmt.Println("2 ------------ Chocolate Chips ($75)")
      fmt.Println("3 ------------ Fruit ($60)")
      fmt.Println("4 ------------ Maraschino Cherry ($30)")
      fmt.Println("5 ------------ Done adding toppings")
      fmt.Print("Select a topping to add (or '5' to finish): ")
      fmt.Scanln(&choice)

      if choice == 5 {
        break
      }
	    var topping IceCream

      switch choice {
      case 1:
        topping = &NutsDecorator{flavor}
      case 2:
        topping = &ChocolateChipsDecorator{flavor}
      case 3:
        topping = &FruitDecorator{flavor}
      case 4:
        topping = &MaraschinoDecorator{flavor}
      default:
        fmt.Println("Invalid topping choice")
        continue
      }

      toppings = append(toppings, topping)
    }

    iceCreamStore.OrderIceCream(flavor, toppings)
    iceCreamRemaining :=iceCreamStore.GetRemainingMoney()
    fmt.Println("Your Remaining Money: ", iceCreamRemaining)
    fmt.Println("Do you want to buy another ice cream?: ")
    

    for {
      fmt.Println("1 - Yes, 2 - No")
      fmt.Scanln(&choice)
    
      if choice == 2 || choice == 1 {
        
        break
      } else {
        fmt.Println("Invalid!")
        continue
      }
    }
    
    
  }
}
	    
