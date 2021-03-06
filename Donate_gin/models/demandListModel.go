package models

import (
	"Donate_gin/dao"
	"strconv"
)

func AddDemanlistModel(Id string,proName string,introduction string,cate string,materials string,recAddress string,cutoff_time string,emergency string)(demandlistId int64,proID int,err error)  {
	recipientId, _ := strconv.Atoi(Id)
	category, _ :=  strconv.Atoi(cate)
	emergencyDegree,_ := strconv.Atoi(emergency)

	demandlistId,err = dao.AddDemandlistDao(recipientId,proName,introduction,category,materials,recAddress,cutoff_time,emergencyDegree)
	if err != nil{
		return
	}

	proID ,err = dao.AddProject(int(demandlistId))
	return
}
