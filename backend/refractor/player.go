package refractor

import (
	"database/sql"
	"github.com/sniddunc/refractor/pkg/broadcast"
)

type Player struct {
	PlayerID      int64    `json:"id"`
	PlayFabID     string   `json:"playFabId"`
	MCUUID        string   `json:"mcuuid"`
	LastSeen      int64    `json:"lastSeen"`
	CurrentName   string   `json:"currentName"`
	PreviousNames []string `json:"previousNames,omitempty"`
}

type DBPlayer struct {
	PlayerID      int64
	PlayFabID     sql.NullString
	MCUUID        sql.NullString
	LastSeen      int64
	CurrentName   string
	PreviousNames []string
}

func (dbp DBPlayer) Player() *Player {
	return &Player{
		PlayerID:      dbp.PlayerID,
		PlayFabID:     dbp.PlayFabID.String,
		MCUUID:        dbp.MCUUID.String,
		LastSeen:      dbp.LastSeen,
		CurrentName:   dbp.CurrentName,
		PreviousNames: dbp.PreviousNames,
	}
}

type PlayerRepository interface {
	Create(player *DBPlayer) error
	FindByID(id int64) (*Player, error)
	FindByPlayFabID(playFabID string) (*Player, error)
	FindByMCUUID(MCUUID string) (*Player, error)
	FindOne(args FindArgs) (*Player, error)
	Exists(args FindArgs) (bool, error)
	UpdateName(player *Player, currentName string) error
	Update(id int64, args UpdateArgs) (*Player, error)
	SearchByName(name string, limit int, offset int) (int, []*Player, error)
}

type PlayerService interface {
	CreatePlayer(newPlayer *DBPlayer) (*Player, *ServiceResponse)
	GetPlayerByID(id int64) (*Player, *ServiceResponse)
	GetPlayer(args FindArgs) (*Player, *ServiceResponse)
	OnPlayerJoin(serverID int64, playerGameID string, currentName string, gameConfig *GameConfig) (*Player, *ServiceResponse)
	OnPlayerQuit(serverID int64, playerGameID string, gameConfig *GameConfig) (*Player, *ServiceResponse)
}

type PlayerHandler interface {
	OnPlayerJoin(fields broadcast.Fields, serverID int64, gameConfig *GameConfig)
	OnPlayerQuit(fields broadcast.Fields, serverID int64, gameConfig *GameConfig)
}
