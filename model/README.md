# model
The module model consist of the logic and templates for creation of all model entities including the type definition for a slice type holding an array of entities. In addition there is the interface which has to implemented by the database access logic in the repository.go

## Entities
Creates all entity models holding all fields with the addition that an ID field will be added by default and in case of lookup field, 1..n child or parent field additional fields will be added. For ever entity a CreatedAt timestamp will be included. This field will be populated on database level when a new record is inserted.

## Repository
Creates a global repository.go file which holds the interfaces to the entity specific repositories

	... interface {
	  Get(id uint64) (*XXX, error)
	  GetAll() (XXXList, error) 
	  Delete(id uint64) error 
	  Update(x *XXX) error 
	  Insert(x *XXX) error 
	  GetLabels() (Labels, error)
	  GetAllXXXByParentID(parentID uint64) (XXXList)				
	}
