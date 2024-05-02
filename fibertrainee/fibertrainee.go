package fibertrainee

import (
	"trainee/fibertrainee/insert"
)

// step1.🍤insert 連線資料庫&放假資料🍤 v
// step2.🍤fiber 設置處理器等待響應http請求 🍤 v
// step3.🍤login 要求登入🍤 v
// step3.🍤token 成功登入後產生jwt token🍤 v
// step4. 🍤塞回cookie🍤
// step5. 🍤登出確認jwt token🍤

type Accontdata struct {
	Account  string
	Password string
}

func FiberTrainee() {
	insert.Insert()
}
