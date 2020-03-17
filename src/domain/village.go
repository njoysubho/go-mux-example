type Village struct{
	gorm.Model
	Name string
}

func (v Village) TotalConsumption() float64{
    return 0.0
}

func ( v Village) create(v Village,db ){

}