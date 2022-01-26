package model

//type DataBaseSessionManager struct{
//	DataBaseSession []*DataBaseSession
//}

//Il pourrait y avoir plusieurs sessions avec un session manager qui serait dans le Global
type DataBaseSession struct {
	Host     string
	Password string
}

//Faire une fonction de connexion qui retoure l'accès au methode de la DB ( mongo / sql / ... )

//func (db *DataBaseSession) Connexion() /* ici retour la struct de la DB */ {
//	fmt.Println(db.Host, db.Password)
//}

//type EntityManager struct {
//	Push func() bool
//	Pull func() bool
//}
