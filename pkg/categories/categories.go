package categories

import "github.com/Yard-Team/whatsapp-api-client-golang/pkg/categories/methods"

type BasisAPICategories struct {
	BasisAPI methods.BasisAPIInterface
}

func (c BasisAPICategories) Account() methods.AccountCategory {
	return methods.AccountCategory{BasisAPI: c.BasisAPI}
}

func (c BasisAPICategories) Device() methods.DeviceCategory {
	return methods.DeviceCategory{BasisAPI: c.BasisAPI}
}

func (c BasisAPICategories) Groups() methods.GroupsCategory {
	return methods.GroupsCategory{BasisAPI: c.BasisAPI}
}

func (c BasisAPICategories) Journals() methods.JournalsCategory {
	return methods.JournalsCategory{BasisAPI: c.BasisAPI}
}

func (c BasisAPICategories) Queues() methods.QueuesCategory {
	return methods.QueuesCategory{BasisAPI: c.BasisAPI}
}

func (c BasisAPICategories) ReadMark() methods.ReadMarkCategory {
	return methods.ReadMarkCategory{BasisAPI: c.BasisAPI}
}

func (c BasisAPICategories) Receiving() methods.ReceivingCategory {
	return methods.ReceivingCategory{BasisAPI: c.BasisAPI}
}

func (c BasisAPICategories) Sending() methods.SendingCategory {
	return methods.SendingCategory{BasisAPI: c.BasisAPI}
}

func (c BasisAPICategories) Service() methods.ServiceCategory {
	return methods.ServiceCategory{BasisAPI: c.BasisAPI}
}
