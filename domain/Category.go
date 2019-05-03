package domain

type Category struct {
	Id string
	Name string
	Site Site
}

func (c Category) GetSite() (Site){
	return c.Site
}
func (c Category) GetIdsConcateneados() (string){
	return c.Site.Id + "-" + c.Id
}

func (c Category) GetAlgo(t string) (string){
	return c.Name + "--" + t
}

func (c Category) CambiarNombre(name string)  {
	c.Name = name

}
func (c *Category) CambiarNombreP(name string)  {
	c.Name = name

}