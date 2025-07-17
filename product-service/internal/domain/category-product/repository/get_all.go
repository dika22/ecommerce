package repository


func (r CategoryProductRepository) GetAll(ctx, dest interface{}) error {
	return r.db.Find(dest).Error	
}