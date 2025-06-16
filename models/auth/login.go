package model_auth

import (
	"strings"
)

func Login(login LoginRequest) error {
	query := `
		select * from (
			select * from (
				select
					div.divoid
				  , div.divname
				  , div.divzip
				  , de.deptoid
				  , de.deptname
				  , lv.leveloid
				  , lv.leveldesc
				  , p.*
				from QL_mstprof p
				inner join QL_mstperson e
					on p.cmpcode = e.cmpcode
						and p.personoid = e.personoid
				inner join QL_mstdivision div
					on div.divoid = e.divisioid
				inner join QL_mstdept de
					on e.deptoid = de.deptoid
				inner join QL_mstlevel lv
					on lv.leveloid = e.leveloid
				union all
				select
					  4 divoid
					, 'PT INTEGRA INDOLESTARI' divname
					, 'IIL' divzip
					, 20 deptoid
					, 'INFORMATION TECHNOLOGY' deptname
					, 0 leveloid
					, '' leveldesc
					, p.*
				from QL_mstprof p where p.profoid = 'admin'
			) as tbl
			where tbl.profapplimit = 1 and tbl.profoid = @username
		) as final
	`

	// jika ada hasil dari query diatas
	if "USER" == "ISEXIST" {
		// jika activeflag == 'ACTIVE'
		if "ACTIVEFLAG" == "ACTIVE" {
			if strings.ToUpper(login.Password) == strings.ToUpper("PROFPASS") {
				query2 := "select top 1 * from itg_profext where profoid = @profoid"
			} else {
				//redirect
				//message = "Password Salah"
			}
		} else {
			// redirect
			// message = "User tidak aktif"
		}
	} else {
		//redirect
		//message = "Username dan Password tidak sama"
	}

}

// yang diambil
// 'profpass'
// $data = [
//     'cmpcode' => $user['cmpcode'],
//     'id' => $user['profoid'],
//     'username' => $user['profoid'],
//     'name' => $user['profname'],
//     'activeflag' => $user['activeflag'],
//     'personoid' => $user['personoid'],
//     'divoid' => $user['divoid'],
//     'divname' => $user['divname'],
//     'divzip' => $user['divzip'],
//     'deptoid' => $user['deptoid'],
//     'deptname' => $user['deptname'],
//     'leveloid' => $user['leveloid'],
//     'leveldesc' => $user['leveldesc'],
//     'role_id' => $role,
// ];
// $this->session->set_userdata($data);
