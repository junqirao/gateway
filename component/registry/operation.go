package registry

type Operation string

const (
	OperationUpdate      Operation = "update"
	OperationDelete      Operation = "delete"
	OperationForceCreate Operation = "force-create"
)

func (o Operation) IsUpdate() bool {
	return o == OperationUpdate
}

func (o Operation) IsDelete() bool {
	return o == OperationDelete
}

func (o Operation) IsForceCreate() bool {
	return o == OperationForceCreate
}

func (o Operation) IsEmpty() bool {
	return o == ""
}
