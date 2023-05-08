package crawlingrules

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-common/util"
	"github.com/rs/zerolog/log"
	"time"
)

type Config struct {
	Name      string        `mapstructure:"name"`
	Disabled  bool          `mapstructure:"disabled"`
	Interval  time.Duration `mapstructure:"interval"`
	Policy    string        `mapstructure:"policy"`
	TimeRange string        `mapstructure:"time-range"`
}

type Rule struct {
	cfg         Config
	validIn     util.TimeOfDayRanges
	lastCheckAt time.Time
	lastCheck   string
	firedAt     time.Time
}

func (dcr *Rule) isCrawlable() bool {

	now := time.Now()
	dcr.lastCheckAt = now

	b := false
	if !dcr.validIn.InRange(now) {
		dcr.lastCheck = "ko-time-range"
	} else {
		switch dcr.cfg.Policy {
		case "once":
			if dcr.firedAt.IsZero() || dcr.lastCheck == "ko-time-range" || util.DayCompare(dcr.firedAt, now) != 0 {
				b = true
			} else {
				dcr.lastCheck = "ko-once"
			}
		case "interval":
			if dcr.firedAt.IsZero() || time.Now().Sub(dcr.firedAt) > dcr.cfg.Interval {
				b = true
			} else {
				dcr.lastCheck = "ko-interval"
			}
		}
	}

	if b {
		dcr.firedAt = now
		dcr.lastCheck = "ok"
	}

	log.Trace().Str("id", dcr.cfg.Name).Str("crawl-able", dcr.lastCheck).Str("cfg", dcr.cfg.Policy).Msg("is crawl-able")
	return b
}

func (dcr *Rule) Crawl() {
	dcr.firedAt = time.Now()
}

type RuleRing struct {
	rules  []Rule
	cursor int // starts at -1
}

func (dcs *RuleRing) current() Rule {
	return dcs.rules[dcs.cursor]
}

// Next returns the index of the selected class.
func (dcs *RuleRing) Next() (Rule, bool) {

	for i := 0; i < len(dcs.rules); i++ {
		dcs.cursor++
		if dcs.cursor == len(dcs.rules) {
			dcs.cursor = 0
		}
		if dcs.rules[dcs.cursor].isCrawlable() {
			return dcs.rules[dcs.cursor], true
		}
	}

	return Rule{}, false
}
