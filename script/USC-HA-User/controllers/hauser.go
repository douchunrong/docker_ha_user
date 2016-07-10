package controllers

import (
	"USC-HA-User/models"
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	//	"gopkg.in/mgo.v2/bson"
)

// operations for Ha_user
type HauserController struct {
	beego.Controller
}

func (c *HauserController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Get", c.Get)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// @Title Post
// @Description create Ha_user
// @Param	body		body 	models.Hauser	true		"body for Ha_user content"
// @Success 201 {object} models.Ha_user
// @Failure 403 body is empty
// @router / [post]
func (c *HauserController) Post() {
	params := make(map[string]interface{})        //接收数据
	callback_data := make(map[string]interface{}) //返回数据

	resq_data := c.Ctx.Input.RequestBody
	fmt.Println("resq_data=======>", string(resq_data))
	json.Unmarshal(resq_data, &params)

	object_id := params["object_id"].(string)
	fmt.Println("=======>", object_id)
	err := models.SaveHaUser(object_id, &params)
	if err != nil {
		callback_data["state"] = 500
	} else {
		callback_data["state"] = 200
	}
	c.Data["json"] = callback_data
	c.ServeJSON()
}

// @Title Get
// @Description get Ha_user by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Ha_user
// @Failure 403 :id is empty
// @router /:id [get]
func (c *HauserController) Get() {
	object_id := c.GetString(":id") //接收数据
	fmt.Println("Get ====>", object_id)
	hauser, err := models.FindHaUser(object_id)
	fmt.Println("hauser ====>", hauser)
	//	callback_data := make(map[string]interface{}) //返回数据
	if err != nil {
		panic(err)
	}
	c.Data["json"] = hauser
	c.ServeJSON()
}

// @Title Get All
// @Description get Ha_user
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Ha_user
// @Failure 403
// @router / [get]
func (c *HauserController) GetAll() {

}

// @Title Update
// @Description update the Ha_user
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Ha_user	true		"body for Ha_user content"
// @Success 200 {object} models.Ha_user
// @Failure 403 :id is not int
// @router /:id [put]
func (c *HauserController) Put() {
	fmt.Println("Put=====>", c.GetString(":id"))
	c.ServeJSON()
}

// @Title Delete
// @Description delete the Ha_user
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *HauserController) Delete() {
	fmt.Println("Delete======>", c.GetString(":id"))
	c.ServeJSON()
}
