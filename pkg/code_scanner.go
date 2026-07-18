package recon
import(

  "io/fs"
  "path/filepath"
  "strings"
)

//CodeScanner yapısı belirli bir dizini tarar

type CodeScanner struct{

  TargetDir string
}

//Scan hedef dizinideki dosyaları gezer

func(s *CodeScanner) Scan(){

  TargetDir string
}

//Scan hedef dizindeki dosyaları gezer
func(s *CodeScanner) Scan(){

  filepath.WalkDir(s.TargetDir, func(path string, d fs.DirEntry, err error) error{

    if err != nil{

      return err
  }

    //Kapsam dışı dizinleri atla (Kurallara uyum [span_0](start_span)[span_0](end_span)

    if strings.Contains(path, "samples") || strings.Contains(path, "tutorials") || string.Contains(path, "demo")
    return nil
}

                   //Sadece .py ve .cs dosyalarını incele 

                   if !id.IsDir() && (strings.HasSuffix(path, ".py") || strings.HasSuffix(path, ".cs")){

                     //Burada dosya içeriğini okuyup analiz edeceğiz 
                     //İleride buraya zafiyet kalıplarımızı (B seçeneği) ekleyeceğiz
                   }
                   return nil
                   })
  }
