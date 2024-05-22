package service

import (
	"context"
	"log"
	"time"
	"tracerstudy-datapipeline-service/common/config"
	"tracerstudy-datapipeline-service/common/utils"
	"tracerstudy-datapipeline-service/modules/pipeline/entity"
	"tracerstudy-datapipeline-service/modules/pipeline/repository"
)

type PipelineService struct {
	cfg                 config.Config
	kabKotaRepository   repository.KabKotaRepositoryUseCase
	provinsiRepository  repository.ProvinsiRepositoryUseCase
	prodiRepository     repository.ProdiRepositoryUseCase
	userStudyRepository repository.UserStudyRepositoryUseCase
	mhsBiodataService   MhsBiodataServiceUseCase
	respondenRepository repository.RespondenRepositoryUseCase
	pktsRepository      repository.PktsRepositoryUseCase
}

type PipelineServiceUseCase interface {
	KabKotaPipeline(ctx context.Context) (uint64, error)
	ProvinsiPipeline(ctx context.Context) (uint64, error)
	ProdiPipeline(ctx context.Context) (uint64, error)
	UserStudyPipeline(ctx context.Context) (uint64, error)
	SiakUpdateRespondenPipeline(ctx context.Context) (uint64, error)
	RespondenPipeline(ctx context.Context) (uint64, error)
	PKTSPipeline(ctx context.Context) (uint64, error)
}

func NewPipelineService(
	cfg config.Config,
	kabKotaRepo repository.KabKotaRepositoryUseCase,
	provinsiRepo repository.ProvinsiRepositoryUseCase,
	prodiRepo repository.ProdiRepositoryUseCase,
	mhsBiodataSvc MhsBiodataServiceUseCase,
	respondenRepo repository.RespondenRepositoryUseCase,
	userStudyRepo repository.UserStudyRepositoryUseCase,
	pktsRepo repository.PktsRepositoryUseCase,
) *PipelineService {
	return &PipelineService{
		cfg:                 cfg,
		kabKotaRepository:   kabKotaRepo,
		provinsiRepository:  provinsiRepo,
		prodiRepository:     prodiRepo,
		userStudyRepository: userStudyRepo,
		mhsBiodataService:   mhsBiodataSvc,
		respondenRepository: respondenRepo,
		pktsRepository:      pktsRepo,
	}
}

func (p *PipelineService) KabKotaPipeline(ctx context.Context) (uint64, error) {
	kabkota, err := p.kabKotaRepository.FindAll(ctx)
	if err != nil {
		return 0, err
	}

	var newKabKota []*entity.NewKabkota
	for _, kk := range kabkota {
		newKabKota = append(newKabKota, &entity.NewKabkota{
			IdWil:      kk.IdWil,
			Nama:       kk.NmWil,
			IdIndukWil: kk.IdIndukWilayah,
			CreatedAt:  kk.CreatedAt,
			UpdatedAt:  kk.UpdatedAt,
		})
	}

	count, err := p.kabKotaRepository.BulkInsert(ctx, newKabKota)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (p *PipelineService) ProvinsiPipeline(ctx context.Context) (uint64, error) {
	provinsi, err := p.provinsiRepository.FindAll(ctx)
	if err != nil {
		return 0, err
	}

	var newProvinsi []*entity.NewProvinsi
	for _, pr := range provinsi {
		newProvinsi = append(newProvinsi, &entity.NewProvinsi{
			IdWil:     pr.IdWil,
			Nama:      pr.NmWil,
			CreatedAt: pr.CreatedAt,
			UpdatedAt: pr.UpdatedAt,
		})
	}

	rows, err := p.provinsiRepository.BulkInsert(ctx, newProvinsi)
	if err != nil {
		return 0, err
	}

	return rows, nil
}

func (p *PipelineService) ProdiPipeline(ctx context.Context) (uint64, error) {
	prodi, err := p.prodiRepository.FindAll(ctx)
	if err != nil {
		return 0, err
	}

	var newProdi []*entity.NewProdi
	for _, pr := range prodi {
		newProdi = append(newProdi, &entity.NewProdi{
			Kode:            pr.Kode,
			KodeDikti:       pr.KodeDikti,
			KodeIntegrasi:   pr.KodeIntegrasi,
			Nama:            pr.Nama,
			Jenjang:         pr.Jenjang,
			KodeFakultas:    pr.KodeFak,
			NamaFakultas:    utils.GetFullNamaFak(pr.NamaFak),
			AkronimFakultas: pr.NamaFak,
			CreatedAt:       pr.CreatedAt,
			UpdatedAt:       pr.UpdatedAt,
		})
	}

	rows, err := p.prodiRepository.BulkInsert(ctx, newProdi)
	if err != nil {
		return 0, err
	}

	return rows, nil
}

func (p *PipelineService) SiakUpdateRespondenPipeline(ctx context.Context) (uint64, error) {
	// err := p.respondenRepository.UpdateStatusUpdate(ctx, "1", "2")
	// if err != nil {
	// 	log.Println("ERROR: [PipelineService - SiakUpdateRespondenPipeline] Error while update status update:", err)
	// 	return err
	// }

	responden, err := p.respondenRepository.FindUnupdated(ctx)
	if err != nil {
		log.Println("ERROR: [PipelineService - SiakUpdateRespondenPipeline] Error while get unupdated responden data:", err)
		return 0, err
	}

	count := uint64(len(responden))

	for _, rs := range responden {
		mhs, err := p.mhsBiodataService.FetchMhsBiodataByNimFromSiakApi(rs.Nim)
		if err != nil {
			count--
			log.Println("ERROR: [PipelineService - SiakUpdateRespondenPipeline] Error while fetch mhs biodata from siak:", err)
			continue
		}

		updateMap := map[string]interface{}{
			"updated_at": time.Now(),
			// "status_update": "1",
			"status_update": "0",
			"ipk":           mhs.IPK,
			"kodedikti":     mhs.KODEPSTD,
			"jenjang":       mhs.JENJANG,
			"namaprodi":     mhs.PRODI,
			"namaprodi2":    mhs.NAMAPST,
			"kodeprodi":     mhs.KODEPST,
			"kodefak":       mhs.KODEFAK,
			"namafak":       mhs.NAMAFAK,
			"jlrmasuk":      mhs.JLRMASUK,
			"thnmasuk":      mhs.THNMASUK,
			"lamastd":       mhs.LAMASTD,
			"tgl_sidang":    mhs.TGLSIDANG,
			"jk":            mhs.KODEJK,
		}

		if mhs.KODEPST != "" && len(mhs.KODEPST) >= 4 {
			updateMap["kodeprodi2"] = mhs.KODEPST[:4]
		}

		if mhs.TGLSIDANG != "" && len(mhs.TGLSIDANG) >= 4 {
			updateMap["thn_sidang"] = mhs.TGLSIDANG[:4]
		}

		if err := p.respondenRepository.Update(ctx, rs.Nim, rs, updateMap); err != nil {
			count--
			log.Println("ERROR: [PipelineService - SiakUpdateRespondenPipeline] Error while update responden data:", err)
			continue
		}

		log.Println("INFO: [PipelineService - SiakUpdateRespondenPipeline] Update responden data success")
	}

	return count, nil
}

func (p *PipelineService) RespondenPipeline(ctx context.Context) (uint64, error) {
	responden, err := p.respondenRepository.FindAll(ctx)
	if err != nil {
		return 0, err
	}

	var newResponden []*entity.NewResponden
	for _, rs := range responden {
		newResponden = append(newResponden, &entity.NewResponden{
			Id:            rs.Id,
			Nim:           rs.Nim,
			Nama:          rs.Nama,
			StatusUpdate:  uint32(utils.ConvStrToUint(rs.StatusUpdate, "status_update")),
			JalurMasuk:    rs.Jlrmasuk,
			TahunMasuk:    rs.Thnmasuk,
			LamaStudi:     uint32(utils.ConvStrToUint(rs.Lamastd, "lama_studi")),
			KodeFakultas:  rs.Kodefak,
			KodeProdi:     rs.Kodeprodi2,
			JenisKelamin:  rs.JK,
			Email:         rs.Email,
			Hp:            utils.FormatPhoneNumber(rs.Hp),
			Ipk:           rs.Ipk,
			TanggalSidang: rs.TglSidang,
			TahunSidang:   rs.ThnSidang,
			TanggalWisuda: rs.TglWisuda,
			Nik:           utils.CleanNumberText(rs.Nik),
			Npwp:          utils.CleanNumberText(rs.Npwp),
			CreatedAt:     rs.CreatedAt,
			UpdatedAt:     rs.UpdatedAt,
		})
	}

	rows, err := p.respondenRepository.BulkInsert(ctx, newResponden)
	if err != nil {
		return 0, err
	}

	return rows, nil
}

func (p *PipelineService) UserStudyPipeline(ctx context.Context) (uint64, error) {
	userStudy, err := p.userStudyRepository.FindAll(ctx)
	if err != nil {
		return 0, err
	}

	var newUserStudy []*entity.NewUserStudy
	for _, us := range userStudy {
		newUserStudy = append(newUserStudy, &entity.NewUserStudy{
			Id:                                uint64(us.ID),
			NamaResponden:                     us.NamaResponden,
			EmailResponden:                    us.EmailResponden,
			HpResponden:                       utils.FormatPhoneNumber(us.HpResponden),
			NamaInstansi:                      us.NamaInstansi,
			Jabatan:                           us.Jabatan,
			AlamatInstansi:                    us.AlamatInstansi,
			NimLulusan:                        us.NimLulusan,
			NamaLulusan:                       us.NamaLulusan,
			ProdiLulusan:                      us.ProdiLulusan,
			TahunLulusan:                      us.TahunLulusan,
			LamaMengenalLulusan:               uint32(utils.ConvStrToUint(us.LamaMengenalLulusan, "LamaMengenalLulusan")),
			Etika:                             us.Etika,
			KeahlianBidIlmu:                   us.KeahlianBidIlmu,
			BahasaInggris:                     us.BahasaInggris,
			PenggunaanTi:                      us.PenggunaanTi,
			Komunikasi:                        us.Komunikasi,
			KerjasamaTim:                      us.KerjasamaTim,
			PengembanganDiri:                  us.PengembanganDiri,
			KesiapanTerjunMasy:                us.KesiapanTerjunMasy,
			KeunggulanLulusan:                 us.KeunggulanLulusan,
			KelemahanLulusan:                  us.KelemahanLulusan,
			SaranPeningkatanKompetensiLulusan: us.SaranPeningkatanKompetensiLulusan,
			SaranPerbaikanKurikulum:           us.SaranPerbaikanKurikulum,
			CreatedAt:                         us.CreatedAt,
			UpdatedAt:                         us.UpdatedAt,
		})
	}

	rows, err := p.userStudyRepository.BulkInsert(ctx, newUserStudy)
	if err != nil {
		return 0, err
	}

	return rows, nil
}

func (p *PipelineService) PKTSPipeline(ctx context.Context) (uint64, error) {
	pkts, err := p.pktsRepository.FindAll(ctx)
	if err != nil {
		return 0, err
	}

	var newPkts []*entity.NewPkts
	for _, pk := range pkts {
		newPkts = append(newPkts, &entity.NewPkts{
			Id:                  uint64(pk.ID),
			Nim:                 pk.Nim,
			KodeProdi:           pk.Kodeprodi,
			TahunSidang:         pk.ThnSidang,
			F8:                  uint16(pk.F8),
			F504:                uint16(pk.F5_04),
			F502:                uint32(utils.CleanNumber(pk.F5_02, "F5_02")),
			F506:                uint32(utils.CleanNumber(pk.F5_06, "F5_06")),
			F505:                uint64(utils.CleanNumber(pk.F5_05, "F5_05")),
			F5a1:                pk.F5a1,
			F5a2:                pk.F5a2,
			F1101:               uint16(pk.F11_01),
			F1102:               pk.F11_02,
			F5b:                 pk.F5b,
			F5c:                 uint16(pk.F5c),
			F5d:                 uint16(pk.F5d),
			F18a:                uint16(pk.F18a),
			F18b:                pk.F18b,
			F18c:                pk.F18c,
			F18d:                pk.F18d,
			F1201:               uint16(pk.F12_01),
			F1202:               pk.F12_02,
			F14:                 uint16(pk.F14),
			F15:                 uint16(pk.F15),
			F1761:               uint16(pk.F1761),
			F1762:               uint16(pk.F1762),
			F1763:               uint16(pk.F1763),
			F1764:               uint16(pk.F1764),
			F1765:               uint16(pk.F1765),
			F1766:               uint16(pk.F1766),
			F1767:               uint16(pk.F1767),
			F1768:               uint16(pk.F1768),
			F1769:               uint16(pk.F1769),
			F1770:               uint16(pk.F1770),
			F1771:               uint16(pk.F1771),
			F1772:               uint16(pk.F1772),
			F1773:               uint16(pk.F1773),
			F1774:               uint16(pk.F1774),
			F21:                 uint16(pk.F21),
			F22:                 uint16(pk.F22),
			F23:                 uint16(pk.F23),
			F24:                 uint16(pk.F24),
			F25:                 uint16(pk.F25),
			F26:                 uint16(pk.F26),
			F27:                 uint16(pk.F27),
			F301:                uint16(pk.F301),
			F302:                uint32(utils.CleanNumber(pk.F302, "F302")),
			F303:                uint32(utils.CleanNumber(pk.F303, "F303")),
			F401:                uint8(utils.CleanNumber(pk.F4_01, "F4_01")),
			F402:                uint8(utils.CleanNumber(pk.F4_02, "F4_02")),
			F403:                uint8(utils.CleanNumber(pk.F4_03, "F4_03")),
			F404:                uint8(utils.CleanNumber(pk.F4_04, "F4_04")),
			F405:                uint8(utils.CleanNumber(pk.F4_05, "F4_05")),
			F406:                uint8(utils.CleanNumber(pk.F4_06, "F4_06")),
			F407:                uint8(utils.CleanNumber(pk.F4_07, "F4_07")),
			F408:                uint8(utils.CleanNumber(pk.F4_08, "F4_08")),
			F409:                uint8(utils.CleanNumber(pk.F4_09, "F4_09")),
			F410:                uint8(utils.CleanNumber(pk.F4_10, "F4_10")),
			F411:                uint8(utils.CleanNumber(pk.F4_11, "F4_11")),
			F412:                uint8(utils.CleanNumber(pk.F4_12, "F4_12")),
			F413:                uint8(utils.CleanNumber(pk.F4_13, "F4_13")),
			F414:                uint8(utils.CleanNumber(pk.F4_14, "F4_14")),
			F415:                uint8(utils.CleanNumber(pk.F4_15, "F4_15")),
			F416:                pk.F4_16,
			F6:                  uint32(utils.CleanNumber(pk.F6, "F6")),
			F7:                  uint32(utils.CleanNumber(pk.F7, "F7")),
			F7a:                 uint32(utils.CleanNumber(pk.F7a, "F7a")),
			F1001:               uint16(pk.F10_01),
			F1002:               pk.F10_02,
			F1601:               uint16(utils.CleanNumber(pk.F16_01, "F16_01")),
			F1602:               uint16(utils.CleanNumber(pk.F16_02, "F16_02")),
			F1603:               uint16(utils.CleanNumber(pk.F16_03, "F16_03")),
			F1604:               uint16(utils.CleanNumber(pk.F16_04, "F16_04")),
			F1605:               uint16(utils.CleanNumber(pk.F16_05, "F16_05")),
			F1606:               uint16(utils.CleanNumber(pk.F16_06, "F16_06")),
			F1607:               uint16(utils.CleanNumber(pk.F16_07, "F16_07")),
			F1608:               uint16(utils.CleanNumber(pk.F16_08, "F16_08")),
			F1609:               uint16(utils.CleanNumber(pk.F16_09, "F16_09")),
			F1610:               uint16(utils.CleanNumber(pk.F16_10, "F16_10")),
			F1611:               uint16(utils.CleanNumber(pk.F16_11, "F16_11")),
			F1612:               uint16(utils.CleanNumber(pk.F16_12, "F16_12")),
			F1613:               uint16(utils.CleanNumber(pk.F16_13, "F16_13")),
			F1614:               pk.F16_14,
			NamaAtasan:          pk.NamaAtasan,
			HpAtasan:            utils.FormatPhoneNumber(pk.HpAtasan),
			EmailAtasan:         pk.EmailAtasan,
			TinggalSelamaKuliah: pk.TinggalSelamaKuliah,
			Code:                pk.Code,
			CreatedAt:           pk.CreatedAt,
			UpdatedAt:           pk.UpdatedAt,
		})
	}

	rows, err := p.pktsRepository.BulkInsert(ctx, newPkts)
	if err != nil {
		return 0, err
	}

	return rows, nil
}
