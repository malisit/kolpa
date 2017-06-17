package kolpa

// Random adress generator function.
// A convenience function, same as g.GenericGenerator("address_address").
// Sample Output: 52185 Katelyn Court Suite 559, Romanstad, AZ 80645
func (g *Generator) Address() string {
	return g.GenericGenerator("address_address")
}

// Random building number generator function.
// A convenience function, same as g.GenericGenerator("address_building_number").
// Sample Output: 27689
func (g *Generator) BuildingNumber() string {
	return g.GenericGenerator("address_building_number")
}

// Random city generator function.
// A convenience function, same as g.GenericGenerator("address_building_number").
// Sample Output: 27689
func (g *Generator) City() string {
	return g.GenericGenerator("address_city")
}

// Random city prefix generator function.
// A convenience function, same as g.GenericGenerator("address_city_prefix").
// Sample Output: Lake
func (g *Generator) CityPrefix() string {
	return g.GenericGenerator("address_city_prefix")
}

// Random city suffix generator function.
// A convenience function, same as g.GenericGenerator("address_city_suffix").
// Sample Output: fort
func (g *Generator) CitySuffix() string {
	return g.GenericGenerator("address_city_suffix")
}

// Random military APO generator function.
// A convenience function, same as g.GenericGenerator("address_military_apo").
// Sample Output: PSC 3455, Box 7670
func (g *Generator) MilitaryAPO() string {
	return g.GenericGenerator("address_military_apo")
}

// Random military DPO generator function.
// A convenience function, same as g.GenericGenerator("address_military_dpo").
// Sample Output: Unit 5426 Box 1476
func (g *Generator) MilitaryDPO() string {
	return g.GenericGenerator("address_military_dpo")
}

// Random military ship generator function.
// A convenience function, same as g.GenericGenerator("address_military_ship_prefix").
// Sample Output: USCGC
func (g *Generator) MilitaryShipPrefix() string {
	return g.GenericGenerator("address_military_ship_prefix")
}

// Random military state abbreviation generator function.
// A convenience function, same as g.GenericGenerator("address_military_state_abbr").
// Sample Output: AE
func (g *Generator) MilitaryStateAbbr() string {
	return g.GenericGenerator("address_military_state_abbr")
}

// Random postcode abbreviation generator function.
// A convenience function, same as g.GenericGenerator("address_postcode").
// Sample Output: 17073
func (g *Generator) Postcode() string {
	return g.GenericGenerator("address_postcode")
}

// Random secondary address generator function.
// A convenience function, same as g.GenericGenerator("address_secondary_address").
// Sample Output: Apt. 786
func (g *Generator) SecondaryAddress() string {
	return g.GenericGenerator("address_secondary_address")
}

// Random state abbreviation generator function.
// A convenience function, same as g.GenericGenerator("address_state_abbr").
// Sample Output: AL
func (g *Generator) StateAbbr() string {
	return g.GenericGenerator("address_state_abbr")
}

// Random state abbreviation generator function.
// A convenience function, same as g.GenericGenerator("address_state").
// Sample Output: Illinois
func (g *Generator) State() string {
	return g.GenericGenerator("address_state")
}

// Random street address generator function.
// A convenience function, same as g.GenericGenerator("address_street_address").
// Sample Output: 98191 Miranda Knoll
func (g *Generator) StreetAddress() string {
	return g.GenericGenerator("address_street_address")
}

// Random street name generator function.
// A convenience function, same as g.GenericGenerator("address_street_name").
// Sample Output: Rodney Station
func (g *Generator) StreetName() string {
	return g.GenericGenerator("address_street_name")
}

// Random street suffix generator function.
// A convenience function, same as g.GenericGenerator("address_street_suffix").
// Sample Output: Pines
func (g *Generator) StreetSuffix() string {
	return g.GenericGenerator("address_street_suffix")
}
