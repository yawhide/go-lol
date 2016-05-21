package lol

import "fmt"

type Match struct {
    Participants []struct {
        Stat struct {
            Winner bool `json:"winner"`
        }   `json:"stats"`
        Timeline struct {
            Role string `json:"role"`
            Lane string `json:"lane"`
        }   `json:"timeline"`
        ParticipantId int `json:"participantId"`
        ChampionId    int `json:"championId"`
        TeamId        int `json:"teamId"`
    }   `json:"participants"`
    ParticipantIdentities []struct {
        Player struct {
            SummonerId uint64 `json:"summonerId"`
        }
        ParticipantId int `json:"participantId"`
    }   `json:"participantIdentities"`
    Teams []struct {
        Winner bool `json:"winner"`
        TeamId int `json:"teamId"`
    }   `json:"teams"`
    MatchId uint64 `json:"matchId"`
}

func (a *APIEndpoint) formatMatchURL(summonerId uint64, options map[string]string) string {
    res := fmt.Sprintf("https://%s/api/lol/%s/v2.2/match/%v?api_key=%s", a.region.url, a.region.code, summonerId, a.key)
    for k, v := range options {
        res = fmt.Sprintf("%s&%s=%s", res, k, v)
    }
    return res
}

func (a *APIEndpoint) GetMatch(matchId uint64, includeTimeline bool) (*Match, error) {
    res := &Match{}
    options := map[string]string{}
    options["includeTimeline"] = fmt.Sprintf("%t", includeTimeline)
    url := a.formatMatchURL(matchId, options)
    err := a.g.Get(url, &res)
    if err != nil {
        return nil, err
    }
    return res, nil
}
