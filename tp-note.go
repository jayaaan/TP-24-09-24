package main

import (
	"fmt"
	"math"
	"sort"
)

func Ft_coin(coins []int, amount int) int {

	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0

	// Parcourir chaque valeur de 1
	for i := 1; i <= amount; i++ {
		// Essayer chaque pièce
		for _, coin := range coins {
			if i-coin >= 0 && dp[i-coin] != math.MaxInt32 {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}

	// Si dp[amount] est toujours infini, cela signifie que la valeur n'est pas atteignable
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Ft_missing(nums []int) int {
	n := len(nums)
	// Calculer la somme des nombres dans l'intervalle [0, n]
	expectedSum := n * (n + 1) / 2
	// Calculer la somme des nombres dans le tableau nums
	actualSum := 0
	for _, num := range nums {
		actualSum += num
	}
	// Le nombre manquant est la différence entre les deux sommes
	return expectedSum - actualSum
}

func Ft_non_overlap(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	// Trier les intervalles par leur fin
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	// Initialiser le compteur d'intervalles non chevauchants et la fin du dernier intervalle sélectionné
	count := 1
	end := intervals[0][1]

	// Parcourir les intervalles et sélectionner ceux qui ne se chevauchent pas
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= end {
			count++
			end = intervals[i][1]
		}
	}

	// Le nombre d'intervalles à retirer est la différence entre le nombre total d'intervalles et le nombre d'intervalles non chevauchants
	return len(intervals) - count
}

func Ft_profit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	// Initialiser le prix minimum et le bénéfice maximum
	minPrice := prices[0]
	maxProfit := 0

	// Parcourir les prix et calculer le bénéfice potentiel
	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else if price-minPrice > maxProfit {
			maxProfit = price - minPrice
		}
	}

	return maxProfit
}

func Ft_max_substring(s string) int {
	if len(s) == 0 {
		return 0
	}

	// Utiliser un tableau pour stocker l'index du dernier caractère vu
	lastIndex := make(map[rune]int)
	maxLength := 0
	start := 0

	// Parcourir la chaîne
	for i, char := range s {
		// Si le caractère est déjà dans la fenêtre, déplacer le début de la fenêtre
		if lastIndex[char] >= start {
			start = lastIndex[char] + 1
		}
		// Mettre à jour l'index du dernier caractère vu
		lastIndex[char] = i
		// Mettre à jour la longueur maximale de la sous-chaîne
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
	}

	return maxLength
}

func Ft_min_window(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}

	// Créer une map pour compter les caractères de t
	dictT := make(map[rune]int)
	for _, char := range t {
		dictT[char]++
	}

	// Créer une map pour compter les caractères dans la fenêtre
	windowCounts := make(map[rune]int)
	have, need := 0, len(dictT)
	result, resultLen := "", math.MaxInt32
	left := 0

	// Parcourir la chaîne s avec une fenêtre glissante
	for right, char := range s {
		windowCounts[char]++
		if dictT[char] > 0 && windowCounts[char] == dictT[char] {
			have++
		}

		// Tant que la fenêtre contient tous les caractères de t, essayer de réduire la fenêtre
		for have == need {
			windowLen := right - left + 1
			if windowLen < resultLen {
				result = s[left : right+1]
				resultLen = windowLen
			}

			// Réduire la fenêtre en déplaçant le pointeur gauche
			windowCounts[rune(s[left])]--
			if dictT[rune(s[left])] > 0 && windowCounts[rune(s[left])] < dictT[rune(s[left])] {
				have--
			}
			left++
		}
	}

	if resultLen == math.MaxInt32 {
		return ""
	}
	return result
}

func main() {
	fmt.Println(Ft_coin([]int{1, 2, 5}, 11)) // résultat : 3 car (11 == 5 + 5 + 1)
	fmt.Println(Ft_coin([]int{2}, 3))        // résultat : -1
	fmt.Println(Ft_coin([]int{1}, 0))        // résultat : 0

	fmt.Println(Ft_missing([]int{3, 1, 2}))                              // résultat : 0
	fmt.Println(Ft_missing([]int{0, 1}))                                 // résultat : 2
	fmt.Println(Ft_missing([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))            // résultat : 8
	fmt.Println(Ft_non_overlap([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}})) // résultat : 1
	fmt.Println(Ft_non_overlap([][]int{{1, 2}, {2, 3}}))                 // résultat : 0
	fmt.Println(Ft_non_overlap([][]int{{1, 2}, {1, 2}, {1, 2}}))         // résultat : 2
	fmt.Println(Ft_profit([]int{7, 1, 5, 3, 6, 4}))                      // résultat : 5
	fmt.Println(Ft_profit([]int{7, 6, 4, 3, 1}))                         // résultat : 0
	fmt.Println(Ft_max_substring("abcabcbb"))                            // résultat : 3
	fmt.Println(Ft_max_substring("bbbbb"))                               // résultat : 1
	fmt.Println(Ft_min_window("ADOBECODEBANC", "ABC"))                   // résultat : "BANC"
	fmt.Println(Ft_min_window("a", "aa"))                                // résultat : ""
}
