package api

import (
	"fmt"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"time"
)

type Coconala struct {
	Name        string        `json:"name"`
	Value       int           `json:"value"`
	URL         string        `json:"url"` // URLを共通化
	WaitTimeSec time.Duration `json:"wait_time_sec"`
}

func NewCoconala() *Coconala {
	return &Coconala{
		URL:         "https://coconala.com", // デフォルトURLを設定
		WaitTimeSec: 1,
	}
}

func (s *Coconala) FetchRootPage() string {
	// 既存のChromiumを使用
	launch := launcher.New().Bin("/usr/bin/chromium").Headless(true).MustLaunch()
	// ブラウザを起動
	browser := rod.New().ControlURL(launch).MustConnect()
	defer browser.MustClose()
	fmt.Println("ブラウザの起動まで完了")

	// ページを開く
	page := browser.MustPage(s.URL)
	fmt.Println("ルートページを開きました")

	// ページの読み込み完了を待つ
	page.MustWaitLoad()

	// ページタイトルを取得
	title := page.MustInfo().Title
	fmt.Printf("Main title: %s\n", title)

	// 「IT・プログラミング・開発」ボタンをクリック
	itCategory := page.MustElementR("span.c-categoryList_listItemName", "IT・プログラミング・開発")
	itCategory.Eval(`this.click()`)

	// ページ遷移を待機
	page.MustWaitLoad()

	fmt.Println("IT・プログラミング・開発に遷移が完了しました。")
	currentURL := page.MustInfo().URL
	fmt.Printf("現在のURL: %s\n", currentURL)
	// カテゴリー名を取得
	page.MustWaitLoad()

	// カテゴリーアイテムを取得
	//要素が表示されるまで待機する
	//time.Sleep(s.WaitTimeSec * time.Second)

	page.MustElement("li.c-searchCategory-item.-child")
	categoryItems := page.MustElements("li.c-searchCategory-item.-child")
	fmt.Printf("カテゴリーアイテム数: %d\n", len(categoryItems))

	// 各カテゴリー名を出力
	for _, item := range categoryItems {
		categoryName := item.MustElement("a").MustText()
		fmt.Println("カテゴリ名:", categoryName)
	}

	// 最初のカテゴリ名を表示
	if len(categoryItems) > 0 {
		// 最初のカテゴリ名を取得
		categoryName := categoryItems[0].MustElement("a").MustText()
		fmt.Println("カテゴリ名:", categoryName)

		// クリックする前に要素が表示され、クリック可能か確認
		// 要素がクリック可能かをチェックする（表示され、クリックできるか）
		// 最初のカテゴリアイテムからリンクを取得
		// 最初のカテゴリをクリックしてリンクを取得
		if len(categoryItems) > 0 {
			aTag := categoryItems[0].MustElement("a")
			href := aTag.MustAttribute("href")
			if href != nil && *href != "" {
				fullURL := s.URL + *href
				fmt.Printf("リンク先: %s\n", fullURL)

				// サブページに遷移
				page.MustNavigate(fullURL)
				page.MustWaitLoad()

				// アイテムリストを取得して処理
				itemLists := page.MustElements(".c-searchPage_itemList .c-searchPageItemList")
				fmt.Printf("c-searchPageItemList の数: %d\n", len(itemLists))

				for _, item := range itemLists {
					// FetchItemDetailsFromLink2 を使用
					s.FetchItemDetailsFromLink2(page, item)
				}
			}
		} else {
			fmt.Println("指定した要素が見つかりませんでした。")
		}
		fmt.Println("Element extraction completed.")

		return title
	}
	return ""
}

// FetchItemDetailsFromLink は item からリンクを取得して、必要な情報を取得する関数
func (s *Coconala) FetchItemDetailsFromLink(page *rod.Page, item *rod.Element) {
	// <a> タグの href 属性を取得して遷移
	aTag := item.MustElement("a")
	href := aTag.MustAttribute("href")

	// href が空でないことを確認
	if href != nil && *href != "" {
		// リンク先に遷移
		fullURL := s.URL + *href // URLを共通化
		page.MustNavigate(fullURL)

		// ページ遷移を待機
		page.MustWaitLoad()

		// 指定した要素がページに現れるまで待つ
		page.MustElement(".c-overview_overview") // ここで待機します

		// タイトルを取得
		titleElement := page.MustElement(".c-overview_overview")
		titleText := titleElement.MustText()
		fmt.Println("タイトル:", titleText)

		// サブタイトルを取得
		subTitleElement := page.MustElement(".c-overview_text")
		subTitleText := subTitleElement.MustText()
		fmt.Println("サブタイトル:", subTitleText)

		// サービス内容を取得
		contentsElement := page.MustElement(".c-contentsFreeText_text")
		contentsText := contentsElement.MustText()
		fmt.Println("サービス内容:", contentsText)

		// 購入にあたってのお願いを取得
		notesElement := page.MustElement(".c-contentsFreeText_text")
		notesText := notesElement.MustText()
		fmt.Println("購入にあたってのお願い:", notesText)

		// 遷移後のURLを取得
		currentURL := page.MustInfo().URL
		fmt.Println("遷移先のURL:", currentURL)
	} else {
		fmt.Println("リンクが見つかりませんでした。")
	}
}

// FetchItemDetailsFromLink2 はアイテムリンクを取得して詳細ページを処理
func (s *Coconala) FetchItemDetailsFromLink2(page *rod.Page, item *rod.Element) {
	aTag := item.MustElement("a")
	href := aTag.MustAttribute("href")
	if href != nil && *href != "" {
		// リンク先に遷移
		fullURL := s.URL + *href
		page.MustNavigate(fullURL)
		page.MustWaitLoad()

		// 必要な情報を取得
		title := page.MustElement(".c-overview_overview").MustText()
		fmt.Println("タイトル:", title)

		subTitle := page.MustElement(".c-overview_text").MustText()
		fmt.Println("サブタイトル:", subTitle)

		contents := page.MustElement(".c-contentsFreeText_text").MustText()
		fmt.Println("サービス内容:", contents)

		notes := page.MustElement(".c-contentsFreeText_text").MustText()
		fmt.Println("購入にあたってのお願い:", notes)

		currentURL := page.MustInfo().URL
		fmt.Println("遷移先のURL:", currentURL)
	} else {
		fmt.Println("リンクが見つかりませんでした。")
	}
}
