package dbservice
	
import (
	"../../models"
)

func (dbcon * DBCON) DeleteGroup(groupName string){
	dbcon.con.Where("group_name = ?", groupName).Delete(&models.Group{})
}
