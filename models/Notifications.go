package models

type Service struct {
	ID 	int64
	Name string
	Price	int64
}

/*
func Payer(connect *pgx.Conn, payer Client, price int64, receiverID int64)  error{
	var balanceAcc int64
	err := connect.QueryRow(context.Background(), `select amount from accounts where Clinetid = $1`, payerID).Scan(&balanceAcc)
	if err != nil {
		fmt.Println("cant get balance %a", err)
		return err
	}
	if price > balanceAcc{
		errors.New("Not enough amount in your balance")
		fmt.Println(err)
		return err
	}
	tx, err := connect.Begin(context.Background())
	if err != nil {
		fmt.Println("cant open transaction %e", err)
	}

	defer func() {
		if err != nil {
			err = tx.Rollback(context.Background())
			if err != nil {
				fmt.Println(err)
			}
		}
		tx.Commit(context.Background())
	}()

	_, err := tx.Exec(`update accounts set amount = $1 where Clientid = $2`, balanceAcc - price, payerID)
	if err != nil {
		return err
	}
	tx.Exec(`update accounts set amount = $1 where Clientid = $2`, price, receiverID)
	return nil

	//записать трансакцию в таблице
}*/