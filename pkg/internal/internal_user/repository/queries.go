package repository

const SelectUser string = "CALL `pssword`.`sp_select_user`(?);"
const SelectUserAndPlan string = "CALL `pssword`.`sp_select_user_and_plan`(?);"
const AddUser string = "CALL `pssword`.`sp_add_user`(?, ?, ?, ?, ?, ?, ?);"
const UpdateUser string = "CALL `pssword`.`sp_update_user`(?, ?, ?, ?, ?, ?, ?);"
const UpdateUserPlan string = "CALL `pssword`.`sp_update_user_plan`(?, ?);"
const UpdateUserPassword string = "CALL `pssword`.`sp_update_user_password`(?, ?, ?);"
const DeleteUser string = "CALL `pssword`.`sp_delete_user`(?, ?);"
