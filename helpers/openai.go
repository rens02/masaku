package helpers

import (
	"context"
	"encoding/json"
	"fmt"
	"masaku/models/web"

	openai "github.com/sashabaranov/go-openai"
	"github.com/sirupsen/logrus"
)

type OpenAiInterface interface {
	GenerateSaran(maag int, asamUrat int, hipertensi int) web.OpenaiRes
}

type OpenAi struct {
	Key string
}

func NewOpenAi(key string) OpenAiInterface {
	return &OpenAi{
		Key: key,
	}
}

func (op *OpenAi) GenerateSaran(maag int, asamUrat int, hipertensi int) web.OpenaiRes {
	var conditions []map[string]interface{}
	var saran = map[string]string{
		"saran_maag":      "Berdasarkan formulir yang kamu isi, ada kemungkinan kamu memiliki kecenderungan maag. Disarankan untuk mengonsumsi makanan yang mudah dicerna dan tidak merangsang asam lambung. Langkah menjaga kesehatan lambung: Hindari makanan pedas, asam, berlemak, dan gorengan. Makan dalam porsi kecil tapi sering, serta hindari makan larut malam. Jangan langsung berbaring setelah makan. Kelola stres, hindari kafein, alkohol, dan rokok. Konsumsi obat lambung sesuai anjuran dokter. Tetap semangat menjaga kesehatan ya, dan konsultasikan ke dokter jika keluhan berlanjut!",
		"saran_asam_urat": "Berdasarkan formulir yang kamu isi, ada kemungkinan kamu memiliki kecenderungan asam urat. Disarankan untuk menghindari makanan tinggi purin seperti jeroan, daging merah, dan makanan laut. Pilih makanan rendah purin seperti sayuran hijau, buah-buahan (pisang, ceri, apel), serta karbohidrat kompleks. Langkah menjaga kesehatan: Perbanyak minum air putih. Hindari alkohol dan minuman manis tinggi fruktosa. Olahraga teratur sesuai kemampuan. Konsumsi obat penurun asam urat sesuai anjuran dokter. Tetap semangat menjaga kesehatan ya, dan konsultasikan ke dokter jika ada keluhan!",
		"saran_hipertensi": "Berdasarkan formulir yang kamu isi, ada kemungkinan kamu memiliki kecenderungan hipertensi. Disarankan untuk mengonsumsi makanan rendah garam, tinggi serat, dan kaya nutrisi. Langkah menjaga tekanan darah: Kurangi garam (maks. 1 sendok teh per hari). Perbanyak sayuran, buah (pisang, alpukat), ikan berlemak, kacang-kacangan, dan oatmeal. Hindari makanan olahan dan tinggi garam. Tetap semangat menjaga kesehatan ya, dan konsultasikan ke dokter jika ada keluhan!",
	}

	// Menambahkan kondisi berdasarkan parameter
	if maag > 0 {
		conditions = append(conditions, map[string]interface{}{
			"condition": "maag",
			"parameter": maag,
		})
	}
	if asamUrat > 0 {
		conditions = append(conditions, map[string]interface{}{
			"condition": "asam urat",
			"parameter": asamUrat,
		})
	}
	if hipertensi > 0 {
		conditions = append(conditions, map[string]interface{}{
			"condition": "hipertensi",
			"parameter": hipertensi,
		})
	}

	// Jika tidak ada kondisi, kembalikan respons default
	if len(conditions) == 0 {
		return web.OpenaiRes{
			Saran: "Hasil formulir menunjukkan tidak ada indikasi signifikan terhadap maag, asam urat, atau hipertensi. Tetap jaga pola hidup sehat dan periksakan diri secara rutin untuk memastikan kesehatanmu tetap optimal!",
		}
	}

	// Mengonversi kondisi menjadi JSON string untuk digunakan dalam prompt
	conditionsJSON, err := json.Marshal(conditions)
	if err != nil {
		logrus.Error("JSON Marshal error: ", err.Error())
		return web.OpenaiRes{
			Saran: "Terjadi kesalahan saat memproses data. Silakan coba lagi nanti.",
		}
	}

	client := openai.NewClient(op.Key)
	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: fmt.Sprintf("Kamu adalah asisten kesehatan yang membantu memberikan saran berdasarkan data kondisi pengguna, buatakan menjadi satu saran yang menyangkut semuanya tanpa perlu level. Contoh saran seperti berikut: %s", saran),
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: fmt.Sprintf("Buatkan rekomendasi kesehatan berdasarkan kondisi berikut: %s", string(conditionsJSON)),
		},
	}

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    openai.GPT3Dot5Turbo,
			Messages: messages,
			MaxTokens: 500,
		},
	)

	if err != nil {
		logrus.Error("CreateChatCompletion error: ", err.Error())
		return web.OpenaiRes{
			Saran: "Terjadi kesalahan saat menghubungi layanan OpenAI. Silakan coba lagi nanti.",
		}
	}

	return web.OpenaiRes{
		Saran: resp.Choices[0].Message.Content,
	}
}
