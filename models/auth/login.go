package model_auth

func Login(login LoginRequest) {
	query := `
	select * from (SELECT div.divoid, div.divname, div.divzip, de.deptoid, de.deptname, lv.leveloid, lv.leveldesc, p.* FROM QL_mstprof p inner join QL_mstperson e on p.cmpcode = e.cmpcode and p.personoid = e.personoid inner join QL_mstdivision div on div.divoid = e.divisioid inner join QL_mstdept de on e.deptoid = de.deptoid inner join QL_mstlevel lv on lv.leveloid = e.leveloid UNION ALL SELECT 4 divoid, 'PT INTEGRA INDOLESTARI' divname, 'IIL' divzip, 20 deptoid, 'INFORMATION TECHNOLOGY' deptname, 0 leveloid, '' leveldesc, p.* FROM QL_mstprof p where p.profoid = 'admin') as tbl where tbl.profapplimit = 1 and tbl.profoid = '$username'
	`
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
