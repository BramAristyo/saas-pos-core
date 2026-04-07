package seeder

import (
	"math/rand"

	"github.com/BramAristyo/saas-pos-core/server/internal/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func SeedProductModifierData(db *gorm.DB) {
	var products []domain.Product
	db.Find(&products)

	var modifierGroups []domain.ModifierGroup
	db.Find(&modifierGroups)

	if len(products) == 0 || len(modifierGroups) == 0 {
		return
	}

	var productModifiers []domain.ProductModifier

	// Create at least 10-15 associations
	// We'll iterate through products and randomly assign 1-3 modifier groups to some of them
	count := 0
	maxAssociations := 15
	if len(products) < 10 {
		maxAssociations = len(products) * 2
	}

	for _, p := range products {
		if count >= maxAssociations {
			break
		}

		// Randomly decide if this product gets modifiers (80% chance for this seeder to ensure enough data)
		if rand.Float32() < 0.8 {
			// Randomly pick 1-2 unique modifier groups for this product
			numModifiers := rand.Intn(2) + 1 // 1 or 2

			// Simple shuffle-based selection to ensure uniqueness for this product
			perm := rand.Perm(len(modifierGroups))
			for i := 0; i < numModifiers && i < len(modifierGroups); i++ {
				mg := modifierGroups[perm[i]]

				productModifiers = append(productModifiers, domain.ProductModifier{
					ProductID:       p.ID,
					ModifierGroupID: mg.ID,
				})
				count++
			}
		}
	}

	// If we still have less than 10 after one pass, just force some more
	if len(productModifiers) < 10 && len(products) > 0 {
		for i := 0; len(productModifiers) < 10 && i < len(products); i++ {
			p := products[i]
			for _, mg := range modifierGroups {
				// Check if already exists in our slice
				exists := false
				for _, existing := range productModifiers {
					if existing.ProductID == p.ID && existing.ModifierGroupID == mg.ID {
						exists = true
						break
					}
				}

				if !exists {
					productModifiers = append(productModifiers, domain.ProductModifier{
						ProductID:       p.ID,
						ModifierGroupID: mg.ID,
					})
					if len(productModifiers) >= 10 {
						break
					}
				}
			}
		}
	}

	db.Clauses(clause.OnConflict{DoNothing: true}).Create(&productModifiers)
}
