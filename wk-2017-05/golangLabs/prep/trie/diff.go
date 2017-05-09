package main

import "fmt"

func getMap(s1 []string) map[string]bool {
  m := make(map[string]bool)
  for i := 0 ; i<len(s1) ; i++ {
    m[s1[i]] = true
  }
  return m
}

func diff(s1, s2 []string) {
  m1 := getMap(s1)
  m2 := getMap(s2)
  
  if len(s1) > len(s2) {
    fmt.Println("> s1 larger")
    for k, _ := range m1 {
      if _, ok := m2[k]; !ok {
        fmt.Println(k)
      }
    }
  } else {
    for k, _ := range m2 {
      if _, ok := m1[k]; !ok {
        fmt.Println(k)
      }
    }
  }
}

func main() {
  out := []string{"gfve","qfpd","lqdqrrcrwdnxeuo","hbsfyfv","ife","feclhbvfuk","ngbqqmbxqcqp","khhqr","dwfcayssyoqc","omwufbdfxu","ilebxwbcto","ta","hbuxhwadlpto","tpvo","phckcyufdqml","lfz","tgygdt","nhcvpf","shwywshtdgmb","bkkxcvg","monmwvytby","qtg","ye","tfsrptug","gama","nberblt","mf","gttxwpuk","xbrtspfttota","qxru","phknqtsdtwxcktmw","pbqurqfxgqlojmws","hdkbdxqg","ge","ukmcowoe","xdnnl","yjyssbsoc","uvouaghhcyvmlk","pbcnmhf","qmmidmvkn","xmywegmtuno","vuzygv","uxtrdsdfzfssmel","eqrswgkaaaohxx","ocedkt","qghoy","wsx","glafbwzdd","ryp","nvybsfrxtlfmp","upmsoxftumyxifyu","xucubv","fctkqlroq","ppvs","nwedtubynsb","repgcx","gsfomhvpmy","kdohe","llkmagl","thkglauzgkeuly","paeurdvexqlw","akdt","rqdll","mumbh","br","fso","qnebmfl","lq","xbhkkfg","ax","gqgsomonv","reqqzzpw","rlbskqgfvn","lfvobeyolyy","mwrvel","ogwilaswn","yw","egdgye","yaqgault","dtlg","iddymnl","evxtehxtbthq","brxpfymuvfvs","rv","udvmara","fecd","dfuydrtbfypbn","cypqodxr","vkmxys","wvpfyfpkvgthq","av","vwduwmjpblqo","xwnbcuggl","flhyfac","mqrbq","pstsxhplrrmbeddv","hnegtuxx","bhhlovgcx","had","aysulvk","potn","os","np","lhv","uzvgyeette","tp","wtsbhgkd","eouxbldsxzm","xhnlcrldtfthul","xhflc","rlaks","phouoyhvls","dqhqats","koat","pybf","po","wygsnxk","kqlima","fgxnuohrnhg","wnnbq","mdtrgv","nkq","agylw","vomtuy","vtdxwrclpspcn","rdrls","yxfeoh","myctacn","fdnor","qfzwuwe","svo","dc","odknlbvxrs","hwc","erpbaxq","rrklkb","wlrwyuy","yyhga","xwdbycdu","htedgvsrhchox","wr","suhesetv","qcagsyqggcf","wljmg","npdbamofynykqv","lmq","oyjmeqvhcrvgm","nxtsnason","gbndakaq","hrtbfnq","fguvomeepxoffg","mat","onlvy","cwcchvsasdylb","dphbfaal","mabkapuoud","vl","ffjfb","svthrfmkoxbho","cvet","ucgqyvopafyttrh","vruh","ukffmudygjavem","dccamymhp","kmm","sc","soqlh","gcdqbcdwbwa","gadgt","pgowefka","cpdxf","sox","fq","lfnrhgsxkhx","loc","fkglorkkvx","ggk","nraoenhd","rrbhfwohfv","yoifoyg","ahmv","byyvhgh","hyw","kedhvwy","rglnpxfqwu","su","mbntqrlwyampdg","jfymrbafmyoc","rhymsno","rkbdlchs","ocbbwwd","exlwulswtvot","tlnc","eo","ua","khxkdxflwxme","kk","cw","pxnktxkm","aa","ngrwqpoh","rwuvd","eruffmlg","bxzovyew","hyfrdngjf","ly","pfgsp","akyprzzphew","ubheeqt","rmbxlcmn","apsbu","khwnykughmwrlk","mvndmua","nlwpw","btgcpqwovwp","sfzdknoxzassc","fg","vhfsf","tyff","blolplosqnfcwx","uwkl","puenodlvotb","naefganqo","cymbrl","wrxw","htknfa","wfrvxqdkhbwwef","vqsghhhutdget","wpccoa","nx","bilt","wqqatye","bwsezd","ww","ss","jmruuqscwjp","nxbk","wd","cfq","gubqavlqffhrzz","caybc","dhaccuualacyl","mtkewbprs","oncggqvr","sqqoffmwkplsgbrp","afumtqugec","nxlbkak","fd","ueolqk","esfmqgvxwfy","npbep","yqssxzsydgllfzmo","tovdtkr","hdykehkefp","ordxzm","dutnbetocxylcey","cr","ngzdc","fxyfqbeoktcc","walsx","brdeumb","dnrrgmem","gvyhnchlimsxc","qe","qm","lt","utqfcqyrrwm","wtelvsqrru","qmqxceuohpiffaq","pmxttqftypfexlv","tg","qa","tccvslp","coqs","oa","lxu","ykbcn","hesvnctfvdsp","ku","at","sxlngbtxmqr","wqrom","krvnaf","hfl","typttkrpfvx","nxqmxr","dhthp","eyelg","npsoqsw","reogbmveofvusdsx","yu","pknt","ckorgrm","bpoaboylced","dmoerc","bhopoqdsref","tmnm","cre","vlrfcsftapyujmw","bcxw","eaum","dybjywyaodsyw","lmu","eocfru","fump","oxpmplpcg","qevdyd","gmuyytguexnyc","lmg","lplrsxznfkoklxlv","twn","bhqultkyfq","saeq","xbuw","kng","uoay","kfykd","armuwp","gtghfxf","pbdrx","adoshnx","rqyex","ng","sid","re","vep","ebwrcpafxzhb","opvrnx","vubuucilxyh","rlcnvnuuqfvhw","goabwrqdoudf","wuznnlyd","vfelxvtggkkk","mxlwbkbklbwfsvr","advraqovan","smkln","kddxywvgqxo","syxngevs","mwo","vg","bpeohsufree","lucidbnlysamvy","urbrmmadea","hghv","gu","uztnpqhdl","rfuyp","xbspa","cnumquohlcgt","tdtuquartspkotm","ugrloq","fkyvqguqq","yy","pwlumocnyuoume","goblttgvmndkqgg","lcblwidhjpu","kbu","pey","bc","lqs","xzluwbtmoxokk","lehdfnr","wruqc","wcbm"}
  expected := []string{"gfve","qfpd","lqdqrrcrwdnxeuo","hbsfyfv","ife","feclhbvfuk","ngbqqmbxqcqp","khhqr","dwfcayssyoqc","omwufbdfxu","ilebxwbcto","ta","hbuxhwadlpto","tpvo","phckcyufdqml","lfz","tgygdt","nhcvpf","shwywshtdgmb","bkkxcvg","monmwvytby","qtg","cwkuzyamnerp","ye","tfsrptug","gama","nberblt","mf","gttxwpuk","xbrtspfttota","qxru","phknqtsdtwxcktmw","pbqurqfxgqlojmws","hdkbdxqg","ge","ukmcowoe","xdnnl","yjyssbsoc","uvouaghhcyvmlk","pbcnmhf","qmmidmvkn","xmywegmtuno","vuzygv","uxtrdsdfzfssmel","eqrswgkaaaohxx","ocedkt","qghoy","wsx","glafbwzdd","ryp","nvybsfrxtlfmp","upmsoxftumyxifyu","xucubv","fctkqlroq","ppvs","nwedtubynsb","repgcx","gsfomhvpmy","kdohe","llkmagl","thkglauzgkeuly","paeurdvexqlw","akdt","rqdll","mumbh","br","fso","qnebmfl","lq","xbhkkfg","ax","gqgsomonv","reqqzzpw","rlbskqgfvn","lfvobeyolyy","mwrvel","ogwilaswn","yw","egdgye","yaqgault","dtlg","iddymnl","evxtehxtbthq","brxpfymuvfvs","rv","udvmara","fecd","dfuydrtbfypbn","cypqodxr","vkmxys","wvpfyfpkvgthq","av","vwduwmjpblqo","xwnbcuggl","flhyfac","mqrbq","pstsxhplrrmbeddv","hnegtuxx","bhhlovgcx","had","aysulvk","potn","os","np","lhv","uzvgyeette","tp","wtsbhgkd","eouxbldsxzm","xhnlcrldtfthul","xhflc","rlaks","phouoyhvls","dqhqats","koat","pybf","po","wygsnxk","kqlima","fgxnuohrnhg","wnnbq","mdtrgv","nkq","agylw","vomtuy","vtdxwrclpspcn","rdrls","yxfeoh","myctacn","fdnor","qfzwuwe","svo","dc","odknlbvxrs","hwc","erpbaxq","rrklkb","wlrwyuy","yyhga","xwdbycdu","htedgvsrhchox","wr","suhesetv","qcagsyqggcf","wljmg","npdbamofynykqv","lmq","oyjmeqvhcrvgm","nxtsnason","gbndakaq","hrtbfnq","fguvomeepxoffg","mat","onlvy","cwcchvsasdylb","dphbfaal","mabkapuoud","vl","ffjfb","svthrfmkoxbho","cvet","ucgqyvopafyttrh","vruh","ukffmudygjavem","dccamymhp","kmm","sc","soqlh","gcdqbcdwbwa","gadgt","pgowefka","cpdxf","sox","fq","lfnrhgsxkhx","loc","fkglorkkvx","ggk","nraoenhd","rrbhfwohfv","yoifoyg","ahmv","byyvhgh","hyw","kedhvwy","rglnpxfqwu","su","mbntqrlwyampdg","jfymrbafmyoc","rhymsno","rkbdlchs","ocbbwwd","exlwulswtvot","tlnc","eo","ua","khxkdxflwxme","kk","cw","pxnktxkm","aa","ngrwqpoh","rwuvd","eruffmlg","bxzovyew","hyfrdngjf","ly","pfgsp","akyprzzphew","ubheeqt","rmbxlcmn","apsbu","khwnykughmwrlk","mvndmua","nlwpw","btgcpqwovwp","sfzdknoxzassc","fg","vhfsf","tyff","blolplosqnfcwx","uwkl","puenodlvotb","naefganqo","cymbrl","wrxw","htknfa","wfrvxqdkhbwwef","vqsghhhutdget","wpccoa","nx","bilt","wqqatye","bwsezd","ww","ss","jmruuqscwjp","nxbk","wd","cfq","gubqavlqffhrzz","caybc","dhaccuualacyl","mtkewbprs","oncggqvr","sqqoffmwkplsgbrp","afumtqugec","nxlbkak","fd","ueolqk","esfmqgvxwfy","npbep","yqssxzsydgllfzmo","tovdtkr","hdykehkefp","ordxzm","dutnbetocxylcey","cr","ngzdc","fxyfqbeoktcc","walsx","brdeumb","dnrrgmem","gvyhnchlimsxc","qe","qm","lt","utqfcqyrrwm","wtelvsqrru","qmqxceuohpiffaq","pmxttqftypfexlv","tg","qa","tccvslp","coqs","oa","lxu","ykbcn","hesvnctfvdsp","ku","at","sxlngbtxmqr","wqrom","krvnaf","hfl","typttkrpfvx","nxqmxr","dhthp","eyelg","npsoqsw","reogbmveofvusdsx","yu","pknt","ckorgrm","bpoaboylced","dmoerc","bhopoqdsref","tmnm","cre","vlrfcsftapyujmw","bcxw","eaum","dybjywyaodsyw","lmu","eocfru","fump","oxpmplpcg","qevdyd","gmuyytguexnyc","lmg","lplrsxznfkoklxlv","twn","bhqultkyfq","saeq","xbuw","kng","uoay","kfykd","armuwp","gtghfxf","pbdrx","adoshnx","rqyex","ng","sid","re","vep","ebwrcpafxzhb","opvrnx","vubuucilxyh","rlcnvnuuqfvhw","goabwrqdoudf","wuznnlyd","vfelxvtggkkk","mxlwbkbklbwfsvr","advraqovan","smkln","kddxywvgqxo","syxngevs","mwo","vg","bpeohsufree","lucidbnlysamvy","urbrmmadea","hghv","gu","uztnpqhdl","rfuyp","xbspa","cnumquohlcgt","tdtuquartspkotm","ugrloq","fkyvqguqq","yy","pwlumocnyuoume","goblttgvmndkqgg","lcblwidhjpu","kbu","pey","bc","lqs","xzluwbtmoxokk","lehdfnr","wruqc","wcbm"}

  fmt.Println("> out", len(out), "> expected >",len(expected))
  // for i := 0 ; i < len(out) || i< len(expected) ; i++ {
  //   if out[i] != expected[i] {
  //     fmt.Println(i, out[i], expected[i])
  //   }
  // }
  diff(out, expected)
}

/*
a simple diffing tool to figure what items in one string list is not apart of the other.
*/