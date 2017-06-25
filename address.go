package kolpa

// Address function returns a random full address.
// A convenience function, same as g.GenericGenerator("address_address").
// Sample Output: 52185 Katelyn Court Suite 559, Romanstad, AZ 80645
func (g *Generator) Address() string {
	return g.GenericGenerator("address_address")
}

// BuildingNumber function returns a random building number.
// A convenience function, same as g.GenericGenerator("address_building_number").
// Sample Output: 27689
func (g *Generator) BuildingNumber() string {
	return g.GenericGenerator("address_building_number")
}

// City function returns a random city name.
// A convenience function, same as g.GenericGenerator("address_building_number").
// Sample Output: 27689
func (g *Generator) City() string {
	return g.GenericGenerator("address_city")
}

// CityPrefix function returns a random city prefix.
// A convenience function, same as g.GenericGenerator("address_city_prefix").
// Sample Output: Lake
func (g *Generator) CityPrefix() string {
	return g.GenericGenerator("address_city_prefix")
}

// CitySuffix function returns a random city suffix.
// A convenience function, same as g.GenericGenerator("address_city_suffix").
// Sample Output: fort
func (g *Generator) CitySuffix() string {
	return g.GenericGenerator("address_city_suffix")
}

// MilitaryAPO function returns a random military APO.
// A convenience function, same as g.GenericGenerator("address_military_apo").
// Sample Output: PSC 3455, Box 7670
func (g *Generator) MilitaryAPO() string {
	return g.GenericGenerator("address_military_apo")
}

// MilitaryDPO function returns a random military DPO.
// A convenience function, same as g.GenericGenerator("address_military_dpo").
// Sample Output: Unit 5426 Box 1476
func (g *Generator) MilitaryDPO() string {
	return g.GenericGenerator("address_military_dpo")
}

// MilitaryShipPrefix function returns a random military ship prefix.
// A convenience function, same as g.GenericGenerator("address_military_ship_prefix").
// Sample Output: USCGC
func (g *Generator) MilitaryShipPrefix() string {
	return g.GenericGenerator("address_military_ship_prefix")
}

// MilitaryStateAbbr function returns a random military state abbreviation.
// A convenience function, same as g.GenericGenerator("address_military_state_abbr").
// Sample Output: AE
func (g *Generator) MilitaryStateAbbr() string {
	return g.GenericGenerator("address_military_state_abbr")
}

// Postcode function returns a random postcode.
// A convenience function, same as g.GenericGenerator("address_postcode").
// Sample Output: 17073
func (g *Generator) Postcode() string {
	return g.GenericGenerator("address_postcode")
}

// SecondaryAddress function returns a random secondary address.
// A convenience function, same as g.GenericGenerator("address_secondary_address").
// Sample Output: Apt. 786
func (g *Generator) SecondaryAddress() string {
	return g.GenericGenerator("address_secondary_address")
}

// StateAbbr function returns a random state name abbreviation.
// A convenience function, same as g.GenericGenerator("address_state_abbr").
// Sample Output: AL
func (g *Generator) StateAbbr() string {
	return g.GenericGenerator("address_state_abbr")
}

// State function returns a random state name.
// A convenience function, same as g.GenericGenerator("address_state").
// Sample Output: Illinois
func (g *Generator) State() string {
	return g.GenericGenerator("address_state")
}

// StreetAddress function returns a random street address.
// A convenience function, same as g.GenericGenerator("address_street_address").
// Sample Output: 98191 Miranda Knoll
func (g *Generator) StreetAddress() string {
	return g.GenericGenerator("address_street_address")
}

// StreetName function returns a random street name.
// A convenience function, same as g.GenericGenerator("address_street_name").
// Sample Output: Rodney Station
func (g *Generator) StreetName() string {
	return g.GenericGenerator("address_street_name")
}

// StreetSuffix function returns a random street suffix.s
// A convenience function, same as g.GenericGenerator("address_street_suffix").
// Sample Output: Pines
func (g *Generator) StreetSuffix() string {
	return g.GenericGenerator("address_street_suffix")
}
