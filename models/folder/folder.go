package models_folder

import (
	"database/sql"
	"file-manager/db"
	model_file "file-manager/models/file"
)

type FolderData struct {
	Title            string                `json:"title"`
	User             string                `json:"user"`
	Lap_harisan_saya []model_file.FileItem `json:"lap_harisan_saya"`
	Folderhidebudept string                `json:"folderhidebudept"`
	Divoid           int                   `json:"divoid"`
	Deptoid          int                   `json:"deptoid"`
	Roleid           int                   `json:"roleid"`
}

func Folder(folderoid, divoid, deptoid int) (FolderData, error) {
	// $this->updlastact();
	// $this->form_validation->set_rules('createuser', 'ID', 'required|trim');

	var result FolderData
	var title, titlehead, foldertype, folderhidebudept, divname, deptname string

	row := db.DB.QueryRow(`
		select top 1 
			  [name]
			, headfolder
			, type
			, folderhidebudept 
		from folder_list where folderoid = @folderoid`,
		sql.Named("folderoid", folderoid),
	)

	err := row.Scan(&title, &titlehead, &foldertype, &folderhidebudept)
	if err != nil {
		return result, err
	}

	if foldertype == "budeptfolder" || foldertype == "bufolder" {
		if divoid != 0 {
			row = db.DB_MIS.QueryRow(`
				select top 1 divname from QL_mstdivision where divoid = @divoid`,
				sql.Named("divoid", divoid),
			)

			err = row.Scan(&divname)
			if err != nil {
				return result, err
			}
		}
		if deptoid != 0 {
			row = db.DB_DEV.QueryRow(`
				select top 1 name from dept_list where divoid=@divoid and deptoid=@deptoid`,
				sql.Named("divoid", divoid),
				sql.Named("deptoid", deptoid),
			)

			err = row.Scan(&deptname)
			if err != nil {
				return result, err
			}
		}
	}

	if foldertype == "subfolder" {
		result.Title = titlehead + " -> " + title
	} else if foldertype == "headfolder" {
		result.Title = title
	} else if foldertype == "budeptfolder" {
		result.Title = title + " -> " + divname + " -> " + deptname
	} else if foldertype == "bufolder" {
		result.Title = title + " -> " + divname
	}

	LapHarian, err := model_file.GetFile(folderoid, divoid, deptoid)
	if err != nil {
		return result, err
	}

	result.User = "admin"               // sementara dummy
	result.Lap_harisan_saya = LapHarian // sementara dummy (diambil dari models file/getLapHarian (folderoid, divoid, deptoid))
	result.Folderhidebudept = folderhidebudept
	result.Divoid = divoid
	result.Deptoid = deptoid
	result.Roleid = 1 // sementara dummy

	return result, nil

	// $upload_file = $_FILES['file']['name'];
	//         if ($upload_file) {
	//             $sql = "";
	//             $fileraw = $this->_prep_filename($_FILES['file']['name']);
	//             $file = preg_replace('/\s+/', '_', $fileraw);
	//             $file = str_replace(",", '', $file);
	//             if ($foldertype == "subfolder") {
	//                 $sql = "select * from itg_file where folderoid = $id and fileurl='$file'";
	//             } else if ($foldertype == "headfolder") {
	//                 $sql = "select * from itg_file where folderoid = $id and fileurl='$file'";
	//             } else if ($foldertype == "budeptfolder") {
	//                 $sql = "select * from itg_file where folderoid = $id and divoid = $divoid and deptoid = $deptoid and fileurl='$file'";
	//             } else if ($foldertype == "bufolder") {
	//                 $sql = "select * from itg_file where folderoid = $id and divoid = $divoid and fileurl='$file'";
	//             }
	//             $valid1 = $this->db->query($sql);
	//             $filecount = $valid1->num_rows();
	//             if ($filecount == 0) {
	//                 $filenumber = $this->input->post('filenumber');
	//                 $sql = "select * from itg_file where folderoid <> 0 and filenumber='$filenumber'";
	//                 $valid2 = $this->db->query($sql);
	//                 $filecount2 = $valid2->num_rows();
	//                 if ($filecount2 == 0) {
	//                     $titlehead = $this->get_laporan->char_replace($titlehead);
	//                     $title = $this->get_laporan->char_replace($title);
	//                     if ($foldertype == "budeptfolder") {
	//                         $divname = $this->get_laporan->char_replace($divname);
	//                         $deptname = $this->get_laporan->char_replace($deptname);
	//                     } else if ($foldertype == "bufolder") {
	//                         $divname = $this->get_laporan->char_replace($divname);
	//                     }
	//                     $config['allowed_types'] = 'xls|xlsx|doc|docx|ppt|pptx|pdf|zip|rar|txt';
	//                     $config['max_size']     = '102400000';
	//                     //$config['upload_path'] = '\\\\fs.integragroup-indonesia.com\\FileManager\\'.$titlehead.'\\'.$title.'\\';
	//                     if ($foldertype == "subfolder") {
	//                         $config['upload_path'] = 'C:\\FileManager\\' . $titlehead . '\\' . $title . '\\';
	//                     } else if ($foldertype == "headfolder") {
	//                         $config['upload_path'] = 'C:\\FileManager\\' . $title . '\\';
	//                     } else if ($foldertype == "budeptfolder") {
	//                         $config['upload_path'] = 'C:\\FileManager\\' . $title . '\\' . $divname . '\\' . $deptname . '\\';
	//                     } else if ($foldertype == "bufolder") {
	//                         $config['upload_path'] = 'C:\\FileManager\\' . $title . '\\' . $divname . '\\';
	//                     }

	//                     //if (!is_dir('\\\\fs.integragroup-indonesia.com\\FileManager\\'.$titlehead.'\\'.$title)) {
	//                     //    mkdir('\\\\fs.integragroup-indonesia.com\\FileManager\\'.$titlehead.'\\'.$title, 0777, TRUE);
	//                     //}
	//                     if ($foldertype == "subfolder") {
	//                         if (!is_dir('C:\\FileManager\\' . $titlehead . '\\' . $title)) {
	//                             mkdir('C:\\FileManager\\' . $titlehead . '\\' . $title, 0777, TRUE);
	//                         }
	//                     } else if ($foldertype == "headfolder") {
	//                         if (!is_dir('C:\\FileManager\\' . $title)) {
	//                             mkdir('C:\\FileManager\\' . $title, 0777, TRUE);
	//                         }
	//                     } else if ($foldertype == "budeptfolder") {
	//                         if (!is_dir('C:\\FileManager\\' . $title . '\\' . $divname . '\\' . $deptname)) {
	//                             mkdir('C:\\FileManager\\' . $title . '\\' . $divname . '\\' . $deptname, 0777, TRUE);
	//                         }
	//                     } else if ($foldertype == "bufolder") {
	//                         if (!is_dir('C:\\FileManager\\' . $title . '\\' . $divname)) {
	//                             mkdir('C:\\FileManager\\' . $title . '\\' . $divname, 0777, TRUE);
	//                         }
	//                     }

	//                     $this->load->library('upload', $config);
	//                     $this->upload->initialize($config);
	//                     if ($this->upload->do_upload('file')) {
	//                         $pos = strrpos($this->upload->data('file_name'), ".pdf");
	//                         if ($pos == true) {
	//                             $src = $this->upload->upload_path . $this->upload->file_name;
	//                             $this->normalizePdf($src);
	//                         }

	//                         $new_file = $this->upload->data('file_name');
	//                         if ($this->input->post('filename') == "") {
	//                             $filenama = $new_file;
	//                         } else {
	//                             $filenama = $this->input->post('filename');
	//                         }
	//                         if ($foldertype == "headfolder") {
	//                             $divoid = $this->session->userdata('divoid');
	//                             $deptoid = $this->session->userdata('deptoid');
	//                         } else if ($foldertype == "subfolder") {
	//                             $divoid = $this->session->userdata('divoid');
	//                             $deptoid = $this->session->userdata('deptoid');
	//                         }
	//                         if ($divzip == "") {
	//                             $divzip = $this->session->userdata('divzip');
	//                         }
	//                         $data = array(
	//                             'divoid' => $divoid,
	//                             'deptoid' => $deptoid,
	//                             'leveloid' => $this->session->userdata('leveloid'),
	//                             'divzip' => $divzip,
	//                             'folderoid' => $id,
	//                             'filename' => $filenama,
	//                             'fileurl' => $new_file,
	//                             'createuser' => $this->input->post('createuser'),
	//                             'createtime' => $this->input->post('createtime'),
	//                             'lastupduser' => $this->input->post('createuser'),
	//                             'lastupdatetime' => $this->input->post('createtime'),
	//                             'filenumber' => $this->input->post('filenumber'),
	//                             'filerevnumber' => $this->input->post('revisionnumber'),
	//                             'filerevdate' => $this->input->post('revisiondate'),
	//                             'fileoldnumber' => '',
	//                             'filevisible' => 'True'
	//                         );
	//                         $this->db->insert('itg_file', $data);
	//                         $this->session->set_flashdata('message', 'Upload data');
	//                         if ($foldertype == "subfolder") {
	//                             redirect('user/folder/' . $id . '/0/0');
	//                         } else if ($foldertype == "headfolder") {
	//                             redirect('user/folder/' . $id . '/0/0');
	//                         } else if ($foldertype == "budeptfolder") {
	//                             redirect('user/folder/' . $id . '/' . $divoid . '/' . $deptoid);
	//                         } else if ($foldertype == "bufolder") {
	//                             redirect('user/folder/' . $id . '/' . $divoid . '/0');
	//                         }
	//                         //redirect('user/folder/'.$id);
	//                     } else {
	//                         $error = $this->upload->display_errors();
	//                         $this->session->set_flashdata('msg', '<div class="alert alert-danger font-weight-bolder" role="alert">' . $error . '</div>');
	//                         if ($foldertype == "subfolder") {
	//                             redirect('user/folder/' . $id . '/0/0');
	//                         } else if ($foldertype == "headfolder") {
	//                             redirect('user/folder/' . $id . '/0/0');
	//                         } else if ($foldertype == "budeptfolder") {
	//                             redirect('user/folder/' . $id . '/' . $divoid . '/' . $deptoid);
	//                         } else if ($foldertype == "bufolder") {
	//                             redirect('user/folder/' . $id . '/' . $divoid . '/0');
	//                         }
	//                         //redirect('user/folder/'.$id);
	//                     }
	//                 } else {
	//                     $this->session->set_flashdata('msg', '<div class="alert alert-danger font-weight-bolder" role="alert">UPLOAD FAILED! file number ' . $filenumber . ' is already exist. please check in All Files</div>');
	//                     if ($foldertype == "subfolder") {
	//                         redirect('user/folder/' . $id . '/0/0');
	//                     } else if ($foldertype == "headfolder") {
	//                         redirect('user/folder/' . $id . '/0/0');
	//                     } else if ($foldertype == "budeptfolder") {
	//                         redirect('user/folder/' . $id . '/' . $divoid . '/' . $deptoid);
	//                     } else if ($foldertype == "bufolder") {
	//                         redirect('user/folder/' . $id . '/' . $divoid . '/0');
	//                     }
	//                 }
	//             } else {
	//                 if ($foldertype == "subfolder") {
	//                     $sql = "select TOP 1 * from itg_file where folderoid = $id and fileurl='$file'";
	//                 } else if ($foldertype == "headfolder") {
	//                     $sql = "select TOP 1 * from itg_file where folderoid = $id and fileurl='$file'";
	//                 } else if ($foldertype == "budeptfolder") {
	//                     $sql = "select TOP 1 * from itg_file where folderoid = $id and divoid = $divoid and deptoid = $deptoid and fileurl='$file'";
	//                 } else if ($foldertype == "bufolder") {
	//                     $sql = "select TOP 1 * from itg_file where folderoid = $id and divoid = $divoid and fileurl='$file'";
	//                 }
	//                 $query = $this->db->query($sql);
	//                 $filenumberold = "";
	//                 foreach ($query->result() as $row) {
	//                     $filenumberold = $row->filenumber;
	//                 }
	//                 $this->session->set_flashdata('msg', '<div class="alert alert-danger font-weight-bolder" role="alert">UPLOAD FAILED! file uploaded is exist in this folder with number ' . $filenumberold . '.</div>');
	//                 if ($foldertype == "subfolder") {
	//                     redirect('user/folder/' . $id . '/0/0');
	//                 } else if ($foldertype == "headfolder") {
	//                     redirect('user/folder/' . $id . '/0/0');
	//                 } else if ($foldertype == "budeptfolder") {
	//                     redirect('user/folder/' . $id . '/' . $divoid . '/' . $deptoid);
	//                 } else if ($foldertype == "bufolder") {
	//                     redirect('user/folder/' . $id . '/' . $divoid . '/0');
	//                 }
	//                 //redirect('user/folder/'.$id);
	//             }
	//         } else {
	//             $this->session->set_flashdata('msg', '<div class="alert alert-danger font-weight-bolder" role="alert">UPLOAD FAILED! No file uploaded.</div>');
	//             if ($foldertype == "subfolder") {
	//                 redirect('user/folder/' . $id . '/0/0');
	//             } else if ($foldertype == "headfolder") {
	//                 redirect('user/folder/' . $id . '/0/0');
	//             } else if ($foldertype == "budeptfolder") {
	//                 redirect('user/folder/' . $id . '/' . $divoid . '/' . $deptoid);
	//             } else if ($foldertype == "bufolder") {
	//                 redirect('user/folder/' . $id . '/' . $divoid . '/0');
	//             }
	//             //redirect('user/folder/'.$id);
	//         }
}
