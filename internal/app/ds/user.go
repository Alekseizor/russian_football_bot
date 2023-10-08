package ds

type User struct {
	VkID               int    `db:"vk_id,omitempty"`
	State              string `db:"state,omitempty"`
	Colors             bool   `db:"colors,omitempty"`
	StyleGame          bool   `db:"style_game,omitempty"`
	History            bool   `db:"history,omitempty"`
	Champions          bool   `db:"champions,omitempty"`
	LocalFans          bool   `db:"local_fans,omitempty"`
	Region             bool   `db:"region,omitempty"`
	StarPlayer         bool   `db:"star_player,omitempty"`
	StrongDef          bool   `db:"strong_def,omitempty"`
	FastGame           bool   `db:"fast_game,omitempty"`
	YoungTalent        bool   `db:"young_talent,omitempty"`
	StrongAttack       bool   `db:"strong_attack,omitempty"`
	ForeignPlayers     bool   `db:"foreign_players,omitempty"`
	TechnicalGame      bool   `db:"technical_game,omitempty"`
	ExperiencedPlayers bool   `db:"experienced_players,omitempty"`
	StrongCharacter    bool   `db:"strong_character,omitempty"`
	YoungCoaches       bool   `db:"young_coaches,omitempty"`
	TeamPlay           bool   `db:"team_play,omitempty"`
	NationwideFans     bool   `db:"nationwide_fans,omitempty"`
	StrongLeadership   bool   `db:"strong_leadership,omitempty"`
	YoungTrainers      bool   `db:"young_trainers,omitempty"`
}

/*//Generate a salted hash for the input string
func (c *Hash) Generate(s string) (string, error) {
	saltedBytes := []byte(s)
	hashedBytes, err := bcrypt.GenerateFromPassword(saltedBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	hash := string(hashedBytes[:])
	return hash, nil
}

//Compare string to generated hash
func (c *Hash) Compare(hash string, s string) error {
	incoming := []byte(s)
	existing := []byte(hash)
	return bcrypt.CompareHashAndPassword(existing, incoming)
}
*/
