package main

import "fmt"

func main() {
	input := "zcncrcnrrlccmhchssgqsqrsstfffnqfnfsswgwjjcmcnnvjvwjvwvfvnvwwhvvwmwhmwwwhbhhldhdmhhdfdbdfbdblllnfllgslllqhlhfhlhqllncnfnsffwmwzzlglzlwlhwwmvmddpvvbrvrdvrdvdhhzdhdpdzzbgbsssrpsrshrshhbqbgbmmtwwcbcrbbvnvhnncscwcwlcwlwnwswbbnwntthzzdlzdzffvqvdqvvtptdptphpddzzvhzvzwwqgqjjjvsspvspsbsffvpvcvjcvjvrvjjgmjjhrrdhhdvhvttjttzptpssnlnjjlnnhmnndjnjwwtjjtbbfflwfwgwvvpjpwpcwcgglmgmdmnmbbqtqctcwcmwwvmvlmvmdvdvhvlhlzhzttghttnvnjntnqqtmtwwhvhrhfhchqchhdrdllnwnfwfjjmsmmnhhshnhpptstjjpnpzzhmmpssgrsrzsrrffzjjhvjhjcjcpcvpccfnfjnfjfllssrdsdnnmffldldppcctcmtctffprpsrrrfhfmflldccnpppvsswppdgpgjpgpqgpqpjppclcjczcnzzzqvzzwlzzqfzffsqqphqhvqqbcccstsdttwcwcvvmjjhqqmllpglgtgwgnnbnrbnnjdndhhbrbvrvwwwblwwwppmdmttztqzqrqjjpggpmgpmpqqmssvnsvnndnvddtztddrrsnsbbshhqmhqmmdnnrsrllqvvfjfpfqpqfqflqlmqmrrqgrqrrhvvfddfhhfnhnlncnzzwwjwswpsscpchcssrzssjqsqrsswbbzggnqnwqqsrqqzbqbddtffcjjdggwlwnwtnnpddcqddsppcwpcwpphvphhhrprqrqttwgwqwppvdpvvrgggmpmddzvvwwthwtttbqtbbftfrffbdfbbspshphpwhppnhnmhhbbshbsbmbdmmtrmrbrprbprphhnmhmvhmhhqqbvqvqcvcwcbwcchbbgwwqffrcffsvvcczrztrzzdhdmhdmmzbbvlvqqtptpvvfsfppdzppmddfvdfdwdfdmdwwlvvqdvqdqmdqqmnmdndsswsfwfvwfwzzhrzhhwfhhpvhppmssnhsnsnwswvwlwdwzwnwswwmqmjmbbdjdbbtllcpllltqlttnvvppznpnttlrrwlwvllphpbblccgjcggbtbwtwdwrrbnndpddbqqnpnrnnqqshspsggtptztpzzclcggtqtgqttcmmbgbfbdfdfsddlrlvrlvvmqqgbqgbqbcbbnpplqqblqbqmbbbdqbbhjjcgcdgdwdjwjqqlsqqnnwffglfglgnlggjgjtgjtgjtgjjclcwcqqvpvwwvgvhgvhvjjzpzbznnvrrrpbpdbdhhmshsnnnwppcbcvvlldccdggqnqgnqqzbzzcfflnnmppcgcmgmrgmrgmgngwwptpzzpfzpfzpzdppgcgtcgtgccvtctzzdmzmlmzlzwllbplblltgtctrrghgvhgvvfvssngsschshrhvrvddghgwwwwzgwgvgccjdjzjwwvdvvqsqshsvvddmddbllvppmlpmlmwlwppjmppwrrdzrrpmrmrdmrmbmzzwttvwwsfswwlflnlrlvvdllzbztbtpbtppnjpnpdnnjrjsshtssvvsjvjfjwffsszvsvjssqzsslpslppjpspqptbncvzrlwtjvsrwtnzzhwdfsmlthvgqgjrpshpbsrrvnsdbqslbcplnpcjqwmwqsnwdcjsdmccbdglwbrcdcqsfhjqhstvhqqdwltqwhhqcrnpvnzjhhbjvqbqhclwggjqfvnfsvcnjjhbmrvbpjqrbljbtltvnsgdfhddlmsdhcrfwvlvbsdrjwjvtnqzhrlqgjmzsmjlpdjsrjmdhmvgwjmfwtqffnzfrtswrlgvvhhqgpzcjwscfqgjmdhtvbgzdlvzfhgqlqbfwsjrmmhrlcwhrcnwwvngcmsrfgczsfqvvmdtmtprfvjrwrwcqwvgmzcjncrzvcswlzsdszvdtwmptnrhgzqwrhjjtbchhpwsdjnqmnsgzwqzvlzlsznpqgvtqnldjqpvndtsjlzhpzsgthbwvwnlbwjlmndqpcdvjdgdzhctpghlfwrtqtvfwdpgrjbmwzqgthjpmlrsqmzsznddhrbjnggqrdntpbngvldnnltfnmdwfhftjvpqbrzqvdzbzzctshzldtcdgfnczglrrjtwswzdvjrfgwztwznbpplmbgwpcmstcsjtqhmzmzsjwsfbjlbnbdtdsmlpdmrrbhdhpzrjdpzhcwsrgfrhmqzqtjfhpvltnthwjrrrpnsbmmwrhsfqbmnvwhpntltsgwgnqhcvlndfrtrfrnlmbhltmtgzhlzqgtsbbnggdjvbslfbczhpghqqcqlrbtpnbqbflfjrmpmwwvjqgvcqtmfggmptqlqstcmdtqlslnnzwbgnstftfsvsjdrmgbzfnzltwbjmqhsvshnmwhftjdndltdpngzwbrjpgpwmqgfsflnhmtzcmdjmzwrsrrrmpvwgggwrhrwtfwdbgbpmwpcdspdtbqvdwwsnwdtrtdtgnfzzsmlbcqdzbsqnrgtvvnlfcdlcgcnfpqbddqcjfqtmpndmnwvfqgjzrltqvlprmbbhmtwbzjjgfhhhfjswpffmjnsdmrcjrwlwpmfrmhljpphlwwwwmgsjcsrcvmfrdjdhshddshpplzsnsphcmdhvllgmdgrvbvjmtpltdthffsvwwhvgqrhmfjfpdswcqldrhpmznffjsntwrnmnpmsshljszbchctptsdlnbcvpfvtlfnzcrljpdwrsjnlpqcpnwvnqhzhqmjbvlbtgslzthlbbjzsgdglbrltzjdshpfbndjtssvsjqlstnrjdzzjvlpqhmwvrsvndcqrqjjcsqvmvrbhngtcfdprlbnqmhqllddgjpdzbjlphntrtgjrdgtbslrtzczbnnlddzzsvqvqvvzjpjqfhztgtsfggdppfdhzsbjzqjmpnmgqzlsdhjjbfpbsbnzpmhwrzjqhczrgcsflfwtrgwbnbrshjpwltntsnsdhmhqlmzdprcrcpcpjnphsmjwhzdqtncdbwgspmnfzsgmpbdhmslqchhhbbwfrghhnfjplsvrtbvplgrwdnbnfsgpwrqczvzlnfsngnsnbwvpmfdcmjztdnrllslnwcfwwnwsvztqmgqtfvmdqrrrmwfmphbcvwwttpmwjjbvqrmlwtwfsjdpcbmdlnlzcqntfzzmslshwprjfhwwpbbdfcdjwllfwcznwpjpwrlsfnnbgzjllrgtzcdcvhdhbtlrcvfdvsdjlzsmwwqvpzfhzjlqpfbqstvfrpcchmtwgbrhqqbglrvzmctdlpnvmglgdtzpbdngtfdnmsmwbgjstzbqwqcdlhfrtqqnhqvpfhdrjqvsvstftdgwwnwpfbfbdcfqnqlwpdnfhhfctwrgdqpbpbmgnfsnbpjfctvdtjnsfqlrtctrnjgltndngcmrdphhsqpjhprbngjzqqhnhhrdwlwwpmhzwshvrtzfgzlrhwghvpvfprbbvflltplpptvrmwcrdqndfqbfqtlqqwvphsmcvnbzghvsptrphhfcgdsslhfbcwhtjcmnpbvqrfgpsjgqpnnwwhjjwqrhhqgznwdzjqbtmmjljjwctqtfgwqbdrjwqwbbcftvjwfdfrgvsrlcrccpvfzdrcjvqfbhddpvrrhjrmhdgchrghbzsqpmgnmslfctblwlvphdfpvtdtwpdfsjwssmgnsvsqpdbqngccsplhmjbwjwtzwsbjhwpwcslqjdchmbvzrbgnwvjrrrdtvhtlzlrbwthzlhhqzzpvpwbzrrbrbtpwnhldhqqltqrqdddfwdmjzgctnlrjrjwvddfmjpnptdmrvnqjvsjfrmlvlqsthhsbvnjlsdzrjngfnqdjfssmvgrchbwmwbbvfqfhvrtwghmrpddnwbrbvbmqvfzbjdsnbzgrtmsfhmsmjtrqsgmpnwwbfwtp"

	N_DISTINCT_REQUIRED := 14

	for i := 0; i < len(input); i++ {
		count := [26]int{}

		for j := 0; j < N_DISTINCT_REQUIRED; j++ {
			count[input[i+j]-'a'] = 1
		}

		sum := 0
		for j := 0; j < len(count); j++ {
			sum += count[j]
		}
		if sum == N_DISTINCT_REQUIRED {
			fmt.Println(i + N_DISTINCT_REQUIRED)
			break
		}
	}
}
