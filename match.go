package lol

import "fmt"

type Match struct {
   MatchDuration uint64 `json:"matchDuration"`
   MatchID uint64 `json:"matchId"`
   MatchMode string `json:"matchMode"`
   MatchType string `json:"matchType"`
   MatchVersion string `josn:"matchVersion"`
   ParticipantIdentities []struct {
     Player struct {
         MatchHistoryUri string `json:"matchHistoryUri"`
         SummonerID uint64 `json:"summonerId"`
         SummonerName string `json:"summonerName"`
     }
     ParticipantID int `json:"participantId"`
   }   `json:"participantIdentities"`
   Participants []struct {
      //TODO
      ChampionID    int `json:"championId"`
      HighestAchievedSeasonTier string `json:"highestAchievedSeasonTier"`
      Masteries []struct {
         MasteryID uint64 `json:"masteryId"`
         Rank uint64 `json:"rank"`
      }
      ParticipantID int `json:"participantId"`
      Runes []struct {
         Rank uint64 `json:"rank"`
         RuneID uint64 `json:"runeId"`
      }
      Spell1ID int `json:"spell1Id"`
      Spell2ID int `json:"spell2Id"`
      ParticipantStats struct {
         Assists uint64 `json:"assists"`
         ChampLevel  uint64 `json:"champLevel"`
         CombatPlayerScore uint64 `json:"combatPlayerScore"`
         Deaths  uint64 `json:"deaths"`
         DoubleKills uint64 `json:"doubleKills"`
         FirstBloodAssist  bool `json:"firstBloodAssist"`
         FirstBloodKill  bool `json:"firstBloodKill"`
         FirstInhibitorAssist  bool `json:"firstInhibitorAssist"`
         FirstInhibitorKill  bool `json:"firstInhibitorKill"`
         FirstTowerAssist  bool `json:"firstTowerAssist"`
         FirstTowerKill  bool `json:"firstTowerKill"`
         GoldEarned  uint64 `json:"goldEarned"`
         GoldSpent uint64 `json:"goldSpent"`
         InhibitorKills  uint64 `json:"inhibitorKills"`
         Item0 uint64 `json:"item0"`
         Item1 uint64 `json:"item1"`
         Item2 uint64 `json:"item2"`
         Item3 uint64 `json:"item3"`
         Item4 uint64 `json:"item4"`
         Item5 uint64 `json:"item5"`
         Item6 uint64 `json:"item6"`
         KillingSprees uint64 `json:"killingSprees"`
         Kills uint64 `json:"kills"`
         LargestCriticalStrike uint64 `json:"largestCriticalStrike"`
         LargestKillingSpree uint64 `json:"largestKillingSpree"`
         LargestMultiKill  uint64 `json:"largestMultiKill"`
         MagicDamageDealt  uint64 `json:"magicDamageDealt"`
         MagicDamageDealtToChampions uint64 `json:"magicDamageDealtToChampions"`
         MagicDamageTaken  uint64 `json:"magicDamageTaken"`
         MinionsKilled uint64 `json:"minionsKilled"`
         NeutralMinionsKilled  uint64 `json:"neutralMinionsKilled"`
         NeutralMinionsKilledEnemyJungle uint64 `json:"neutralMinionsKilledEnemyJungle"`
         NeutralMinionsKilledTeamJungle  uint64 `json:"neutralMinionsKilledTeamJungle"`
         NodeCapture uint64 `json:"nodeCapture"`
         NodeCaptureAssist uint64 `json:"nodeCaptureAssist"`
         NodeNeutralize  uint64 `json:"nodeNeutralize"`
         NodeNeutralizeAssist  uint64 `json:"nodeNeutralizeAssist"`
         ObjectivePlayerScore  uint64 `json:"objectivePlayerScore"`
         PentaKills  uint64 `json:"pentaKills"`
         PhysicalDamageDealt uint64 `json:"physicalDamageDealt"`
         PhysicalDamageDealtToChampions  uint64 `json:"physicalDamageDealtToChampions"`
         PhysicalDamageTaken uint64 `json:"physicalDamageTaken"`
         QuadraKills uint64 `json:"quadraKills"`
         SightWardsBoughtInGame  uint64 `json:"sightWardsBoughtInGame"`
         TeamObjective uint64 `json:"teamObjective"`
         TotalDamageDealt  uint64 `json:"totalDamageDealt"`
         TotalDamageDealtToChampions uint64 `json:"totalDamageDealtToChampions"`
         TotalDamageTaken  uint64 `json:"totalDamageTaken"`
         TotalHeal uint64 `json:"totalHeal"`
         TotalPlayerScore  uint64 `json:"totalPlayerScore"`
         TotalScoreRank  uint64 `json:"totalScoreRank"`
         TotalTimeCrowdControlDealt  uint64 `json:"totalTimeCrowdControlDealt"`
         TotalUnitsHealed  uint64 `json:"totalUnitsHealed"`
         TowerKills  uint64 `json:"towerKills"`
         TripleKills uint64 `json:"tripleKills"`
         TrueDamageDealt uint64 `json:"trueDamageDealt"`
         TrueDamageDealtToChampions  uint64 `json:"trueDamageDealtToChampions"`
         TrueDamageTaken uint64 `json:"trueDamageTaken"`
         UnrealKills uint64 `json:"unrealKills"`
         VisionWardsBoughtInGame uint64 `json:"visionWardsBoughtInGame"`
         wardsKilled uint64 `json:"wardsKilled"`
         WardsPlaced uint64 `json:"wardsPlaced"`
         Winner bool `json:"winner"`
      }   `json:"stats"`
      TeamID        int `json:"teamId"`
      ParticipantTimeline struct {
         AncientGolemAssistsPerMinCounts struct {
            TenToTwenty float64   `json:"tenToTwenty"`
            ThirtyToEnd float64   `json:"thirtyToEnd"`
            TwentyToThirty  float64   `json:"twentyToThirty"`
            ZeroToTen float64 `json:"zeroToTen"`
         } `json:"ancientGolemAssistsPerMinCounts"`
         AncientGolemKillsPerMinCounts struct {
            TenToTwenty float64   `json:"tenToTwenty"`
            ThirtyToEnd float64   `json:"thirtyToEnd"`
            TwentyToThirty  float64   `json:"twentyToThirty"`
            ZeroToTen float64 `json:"zeroToTen"`
         } `json:"ancientGolemKillsPerMinCounts"`
         AssistedLaneDeathsPerMinDeltas  struct {
            TenToTwenty float64   `json:"tenToTwenty"`
            ThirtyToEnd float64   `json:"thirtyToEnd"`
            TwentyToThirty  float64   `json:"twentyToThirty"`
            ZeroToTen float64 `json:"zeroToTen"`
         } `json:"assistedLaneDeathsPerMinDeltas"`
         AssistedLaneKillsPerMinDeltas struct {
            TenToTwenty float64   `json:"tenToTwenty"`
            ThirtyToEnd float64   `json:"thirtyToEnd"`
            TwentyToThirty  float64   `json:"twentyToThirty"`
            ZeroToTen float64 `json:"zeroToTen"`
         } `json:"assistedLaneKillsPerMinDeltas"`
         BaronAssistsPerMinCounts  struct {
            TenToTwenty float64   `json:"tenToTwenty"`
            ThirtyToEnd float64   `json:"thirtyToEnd"`
            TwentyToThirty  float64   `json:"twentyToThirty"`
            ZeroToTen float64 `json:"zeroToTen"`
         } `json:"baronAssistsPerMinCounts"`
         BaronKillsPerMinCounts  struct {
            TenToTwenty float64   `json:"tenToTwenty"`
            ThirtyToEnd float64   `json:"thirtyToEnd"`
            TwentyToThirty  float64   `json:"twentyToThirty"`
            ZeroToTen float64 `json:"zeroToTen"`
         } `json:"baronKillsPerMinCounts"`
         CreepsPerMinDeltas  struct {
            TenToTwenty float64   `json:"tenToTwenty"`
            ThirtyToEnd float64   `json:"thirtyToEnd"`
            TwentyToThirty  float64   `json:"twentyToThirty"`
            ZeroToTen float64 `json:"zeroToTen"`
         } `json:"creepsPerMinDeltas"`
         CsDiffPerMinDeltas  struct {
            TenToTwenty float64   `json:"tenToTwenty"`
            ThirtyToEnd float64   `json:"thirtyToEnd"`
            TwentyToThirty  float64   `json:"twentyToThirty"`
            ZeroToTen float64 `json:"zeroToTen"`
         } `json:"csDiffPerMinDeltas"`
         DamageTakenDiffPerMinDeltas struct {
            TenToTwenty float64   `json:"tenToTwenty"`
            ThirtyToEnd float64   `json:"thirtyToEnd"`
            TwentyToThirty  float64   `json:"twentyToThirty"`
            ZeroToTen float64 `json:"zeroToTen"`
         } `json:"damageTakenDiffPerMinDeltas"`
         DamageTakenPerMinDeltas struct {
            TenToTwenty float64   `json:"tenToTwenty"`
            ThirtyToEnd float64   `json:"thirtyToEnd"`
            TwentyToThirty  float64   `json:"twentyToThirty"`
            ZeroToTen float64 `json:"zeroToTen"`
         } `json:"damageTakenPerMinDeltas"`
         DragonAssistsPerMinCounts struct {
            TenToTwenty float64   `json:"tenToTwenty"`
            ThirtyToEnd float64   `json:"thirtyToEnd"`
            TwentyToThirty  float64   `json:"twentyToThirty"`
            ZeroToTen float64 `json:"zeroToTen"`
         } `json:"dragonAssistsPerMinCounts"`
         DragonKillsPerMinCounts struct {
            TenToTwenty float64   `json:"tenToTwenty"`
            ThirtyToEnd float64   `json:"thirtyToEnd"`
            TwentyToThirty  float64   `json:"twentyToThirty"`
            ZeroToTen float64 `json:"zeroToTen"`
         } `json:"dragonKillsPerMinCounts"`
         ElderLizardAssistsPerMinCounts  struct {
            TenToTwenty float64   `json:"tenToTwenty"`
            ThirtyToEnd float64   `json:"thirtyToEnd"`
            TwentyToThirty  float64   `json:"twentyToThirty"`
            ZeroToTen float64 `json:"zeroToTen"`
         } `json:"elderLizardAssistsPerMinCounts"`
         ElderLizardKillsPerMinCounts  struct {
            TenToTwenty float64   `json:"tenToTwenty"`
            ThirtyToEnd float64   `json:"thirtyToEnd"`
            TwentyToThirty  float64   `json:"twentyToThirty"`
            ZeroToTen float64 `json:"zeroToTen"`
         } `json:"elderLizardKillsPerMinCounts"`
         GoldPerMinDeltas  struct {
            TenToTwenty float64   `json:"tenToTwenty"`
            ThirtyToEnd float64   `json:"thirtyToEnd"`
            TwentyToThirty  float64   `json:"twentyToThirty"`
            ZeroToTen float64 `json:"zeroToTen"`
         } `json:"goldPerMinDeltas"`
         InhibitorAssistsPerMinCounts  struct {
            TenToTwenty float64   `json:"tenToTwenty"`
            ThirtyToEnd float64   `json:"thirtyToEnd"`
            TwentyToThirty  float64   `json:"twentyToThirty"`
            ZeroToTen float64 `json:"zeroToTen"`
         } `json:"inhibitorAssistsPerMinCounts"`
         InhibitorKillsPerMinCounts  struct {
            TenToTwenty float64   `json:"tenToTwenty"`
            ThirtyToEnd float64   `json:"thirtyToEnd"`
            TwentyToThirty  float64   `json:"twentyToThirty"`
            ZeroToTen float64 `json:"zeroToTen"`
         } `json:"inhibitorKillsPerMinCounts"`
         Lane  string `json:"lane"`
         Role  string `json:"role"`
         TowerAssistsPerMinCounts  struct {
            TenToTwenty float64   `json:"tenToTwenty"`
            ThirtyToEnd float64   `json:"thirtyToEnd"`
            TwentyToThirty  float64   `json:"twentyToThirty"`
            ZeroToTen float64 `json:"zeroToTen"`
         } `json:"towerAssistsPerMinCounts"`
         TowerKillsPerMinCounts  struct {
            TenToTwenty float64   `json:"tenToTwenty"`
            ThirtyToEnd float64   `json:"thirtyToEnd"`
            TwentyToThirty  float64   `json:"twentyToThirty"`
            ZeroToTen float64 `json:"zeroToTen"`
         } `json:"towerKillsPerMinCounts"`
         TowerKillsPerMinDeltas  struct {
            TenToTwenty float64   `json:"tenToTwenty"`
            ThirtyToEnd float64   `json:"thirtyToEnd"`
            TwentyToThirty  float64   `json:"twentyToThirty"`
            ZeroToTen float64 `json:"zeroToTen"`
         } `json:"towerKillsPerMinDeltas"`
         VilemawAssistsPerMinCounts  struct {
            TenToTwenty float64   `json:"tenToTwenty"`
            ThirtyToEnd float64   `json:"thirtyToEnd"`
            TwentyToThirty  float64   `json:"twentyToThirty"`
            ZeroToTen float64 `json:"zeroToTen"`
         } `json:"vilemawAssistsPerMinCounts"`
         VilemawKillsPerMinCounts  struct {
            TenToTwenty float64   `json:"tenToTwenty"`
            ThirtyToEnd float64   `json:"thirtyToEnd"`
            TwentyToThirty  float64   `json:"twentyToThirty"`
            ZeroToTen float64 `json:"zeroToTen"`
         } `json:"vilemawKillsPerMinCounts"`
         WardsPerMinDeltas struct {
            TenToTwenty float64   `json:"tenToTwenty"`
            ThirtyToEnd float64   `json:"thirtyToEnd"`
            TwentyToThirty  float64   `json:"twentyToThirty"`
            ZeroToTen float64 `json:"zeroToTen"`
         } `json:"wardsPerMinDeltas"`
         XpDiffPerMinDeltas  struct {
            TenToTwenty float64   `json:"tenToTwenty"`
            ThirtyToEnd float64   `json:"thirtyToEnd"`
            TwentyToThirty  float64   `json:"twentyToThirty"`
            ZeroToTen float64 `json:"zeroToTen"`
         } `json:"xpDiffPerMinDeltas"`
         XpPerMinDeltas  struct {
            TenToTwenty float64   `json:"tenToTwenty"`
            ThirtyToEnd float64   `json:"thirtyToEnd"`
            TwentyToThirty  float64   `json:"twentyToThirty"`
            ZeroToTen float64 `json:"zeroToTen"`
         } `json:"xpPerMinDeltas"`
      }   `json:"timeline"`
   }   `json:"participants"`
   PlatformID string `json:"platformId"`
   QueueType string `json:"queueType"`
   Region string `json:"region"`
   Season string `json:"season"`
   Teams []struct {
      Bans  []struct {
         ChampionID int `json:"championId"`
         PickTurn int `json:"pickTurn"`
      } `json:"bans"`
      BaronKills  int `json:"baronKills"`
      DominionVictoryScore  uint64 `json:"dominionVictoryScore"`
      DragonKills int `json:"dragonKills"`
      FirstBaron  bool `json:"firstBaron"`
      FirstBlood  bool `json:"firstBlood"`
      FirstDragon bool `json:"firstDragon"`
      FirstInhibitor  bool `json:"firstInhibitor"`
      FirstRiftHerald bool `json:"firstRiftHerald"`
      FirstTower  bool `json:"firstTower"`
      InhibitorKills  int `json:"inhibitorKills"`
      RiftHeraldKills int `json:"riftHeraldKills"`
      TeamID  int `json:"teamId"`
      TowerKills  int `json:"towerKills"`
      VilemawKills  int `json:"vilemawKills"`
      Winner  bool `json:"winner"`
   }   `json:"teams"`

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
