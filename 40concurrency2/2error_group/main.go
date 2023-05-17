package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"

	"golang.org/x/sync/errgroup"
)

type UserType int8

const (
	// tell why not iota
	TypeInvalid      UserType = 0
	TypeRegular      UserType = 1
	TypePremium      UserType = 2
	TypeExperimental UserType = 3
)

type Service struct {
}

type GetUserDataResp struct {
	Users []GetUserDataDto
}
type GetUserDataDto struct {
	ID           int64
	Type         UserType
	BalancePenny int64
}

func (dto GetUserDataDto) String() string {
	return fmt.Sprintf("ID of user: %d, Type of user: %d, Balance: %d", dto.ID, dto.Type, dto.BalancePenny)
}

type UserEntity struct {
	UserID       int64
	Type         UserType
	BalancePenny int64
}

func (entity UserEntity) String() string {
	return fmt.Sprintf("ID of user: %d, Type of user: %d, Balance: %d", entity.UserID, entity.Type, entity.BalancePenny)
}

type Service2 struct {
}

type CanTakeComissionReq struct {
	userIDs    []int64
	throwError bool
}

type CanTakeComissionResp struct {
	Users []int64
}

func main() {
	rand.Seed(time.Now().Unix())
	srv1 := Service{}
	srv2 := Service2{}

	ctx := context.Background()
	g, ctx := errgroup.WithContext(ctx)

	userLen := 10
	userIDs := make([]int64, userLen)
	for i := 0; i < userLen; i++ {
		userIDs[i] = int64(i + 1)
	}

	var (
		userDataResp      *GetUserDataResp
		userCommissionMap map[int64]struct{}
	)

	// can you semaphore
	g.SetLimit(5)

	g.Go(func() error {
		resp, err := srv1.GetUserData(userIDs)
		if err != nil {
			return fmt.Errorf("srv1.GetUserData: %w", err)
		}

		userDataResp = resp

		return nil
	})

	g.Go(func() error {
		resp, err := srv2.CanTakeCommission(&CanTakeComissionReq{
			userIDs:    userIDs,
			throwError: true,
		})
		if err != nil {
			return fmt.Errorf("srv2.CanTakeCommission: %s", err.Error())
		}

		userCommissionMap = make(map[int64]struct{}, len(resp.Users))
		for _, user := range resp.Users {
			userCommissionMap[user] = struct{}{}
		}

		return nil
	})

	if err := g.Wait(); err != nil {
		log.Fatalf("Fatal err from errG: %s", err.Error())
	}

	fmt.Println("Successful result")
	fmt.Println()

	for _, user := range userDataResp.Users {
		fmt.Println(user)

	}
	fmt.Println()

	for userID := range userCommissionMap {
		fmt.Printf("users with comission: %d\n", userID)
	}

	finalResult := mapResult(userDataResp.Users, userCommissionMap)
	fmt.Println()

	fmt.Println(finalResult)
}

func (s Service) GetUserData(userIDs []int64) (*GetUserDataResp, error) {
	fmt.Println("GetUserData start ...")

	// go to some service
	users := make([]GetUserDataDto, 0, len(userIDs))
	for _, ID := range userIDs {
		users = append(users, GetUserDataDto{
			ID:   ID,
			Type: UserType(0 + rand.Intn(4)),
			// Penny
			BalancePenny: rand.Int63(),
		})
	}

	return &GetUserDataResp{
		Users: users,
	}, nil
}

func (s Service2) CanTakeCommission(in *CanTakeComissionReq) (*CanTakeComissionResp, error) {
	fmt.Println("CanTakeCommission start ...")

	// validation
	if in == nil {
		return nil, errors.New("nil in")
	}

	if len(in.userIDs) == 0 {
		return nil, errors.New("nil users")
	}

	if in.throwError {
		return nil, errors.New("throw err for test")
	}

	users := make([]int64, 0, len(in.userIDs))
	for _, ID := range in.userIDs {
		if rand.Intn(2)%2 == 0 {
			users = append(users, ID)
		}
	}

	return &CanTakeComissionResp{
		Users: users,
	}, nil
}

func mapResult(userData []GetUserDataDto, userComissionMap map[int64]struct{}) []UserEntity {
	if len(userData) == 0 || len(userComissionMap) == 0 {
		return nil
	}

	result := make([]UserEntity, 0, len(userComissionMap))
	for _, user := range userData {
		if user.Type == TypeInvalid {
			continue
		}
		if _, ok := userComissionMap[user.ID]; !ok {
			continue
		}

		result = append(result, UserEntity{
			UserID:       user.ID,
			Type:         user.Type,
			BalancePenny: user.BalancePenny,
		})
	}

	return result
}
