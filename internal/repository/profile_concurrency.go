package repository

type CompetencyProfile struct { version int }
func (p *CompetencyProfile) Advance() {
 current := p.version
 p.version = current + 1
}
func (p *CompetencyProfile) Version() int { return p.version }
