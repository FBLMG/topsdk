package domain


type TaobaoJuItemsSearchItems struct {
    /*
        卖点描述     */
    UspDescList  *[]string `json:"usp_desc_list,omitempty" `

    /*
        淘宝类目id     */
    TbFirstCatId  *int64 `json:"tb_first_cat_id,omitempty" `

    /*
        原价     */
    OrigPrice  *string `json:"orig_price,omitempty" `

    /*
        itemId     */
    ItemId  *int64 `json:"item_id,omitempty" `

    /*
        展示结束时间     */
    ShowEndTime  *int64 `json:"show_end_time,omitempty" `

    /*
        pc链接     */
    PcUrl  *string `json:"pc_url,omitempty" `

    /*
        频道id     */
    PlatformId  *int64 `json:"platform_id,omitempty" `

    /*
        聚划算id     */
    JuId  *int64 `json:"ju_id,omitempty" `

    /*
        无线主图     */
    PicUrlForWL  *string `json:"pic_url_for_w_l,omitempty" `

    /*
        开团时间     */
    OnlineStartTime  *int64 `json:"online_start_time,omitempty" `

    /*
        类目名称     */
    CategoryName  *string `json:"category_name,omitempty" `

    /*
        聚划算价格，单位分     */
    ActPrice  *string `json:"act_price,omitempty" `

    /*
        商品标题     */
    Title  *string `json:"title,omitempty" `

    /*
        无线链接     */
    WapUrl  *string `json:"wap_url,omitempty" `

    /*
        商品卖点     */
    ItemUspList  *[]string `json:"item_usp_list,omitempty" `

    /*
        开始展示时间     */
    ShowStartTime  *int64 `json:"show_start_time,omitempty" `

    /*
        开团结束时间     */
    OnlineEndTime  *int64 `json:"online_end_time,omitempty" `

    /*
        pc主图     */
    PicUrlForPC  *string `json:"pic_url_for_p_c,omitempty" `

    /*
        价格卖点     */
    PriceUspList  *[]string `json:"price_usp_list,omitempty" `

    /*
        是否包邮     */
    PayPostage  *bool `json:"pay_postage,omitempty" `

}

func (s *TaobaoJuItemsSearchItems) SetUspDescList(v []string) *TaobaoJuItemsSearchItems {
    s.UspDescList = &v
    return s
}
func (s *TaobaoJuItemsSearchItems) SetTbFirstCatId(v int64) *TaobaoJuItemsSearchItems {
    s.TbFirstCatId = &v
    return s
}
func (s *TaobaoJuItemsSearchItems) SetOrigPrice(v string) *TaobaoJuItemsSearchItems {
    s.OrigPrice = &v
    return s
}
func (s *TaobaoJuItemsSearchItems) SetItemId(v int64) *TaobaoJuItemsSearchItems {
    s.ItemId = &v
    return s
}
func (s *TaobaoJuItemsSearchItems) SetShowEndTime(v int64) *TaobaoJuItemsSearchItems {
    s.ShowEndTime = &v
    return s
}
func (s *TaobaoJuItemsSearchItems) SetPcUrl(v string) *TaobaoJuItemsSearchItems {
    s.PcUrl = &v
    return s
}
func (s *TaobaoJuItemsSearchItems) SetPlatformId(v int64) *TaobaoJuItemsSearchItems {
    s.PlatformId = &v
    return s
}
func (s *TaobaoJuItemsSearchItems) SetJuId(v int64) *TaobaoJuItemsSearchItems {
    s.JuId = &v
    return s
}
func (s *TaobaoJuItemsSearchItems) SetPicUrlForWL(v string) *TaobaoJuItemsSearchItems {
    s.PicUrlForWL = &v
    return s
}
func (s *TaobaoJuItemsSearchItems) SetOnlineStartTime(v int64) *TaobaoJuItemsSearchItems {
    s.OnlineStartTime = &v
    return s
}
func (s *TaobaoJuItemsSearchItems) SetCategoryName(v string) *TaobaoJuItemsSearchItems {
    s.CategoryName = &v
    return s
}
func (s *TaobaoJuItemsSearchItems) SetActPrice(v string) *TaobaoJuItemsSearchItems {
    s.ActPrice = &v
    return s
}
func (s *TaobaoJuItemsSearchItems) SetTitle(v string) *TaobaoJuItemsSearchItems {
    s.Title = &v
    return s
}
func (s *TaobaoJuItemsSearchItems) SetWapUrl(v string) *TaobaoJuItemsSearchItems {
    s.WapUrl = &v
    return s
}
func (s *TaobaoJuItemsSearchItems) SetItemUspList(v []string) *TaobaoJuItemsSearchItems {
    s.ItemUspList = &v
    return s
}
func (s *TaobaoJuItemsSearchItems) SetShowStartTime(v int64) *TaobaoJuItemsSearchItems {
    s.ShowStartTime = &v
    return s
}
func (s *TaobaoJuItemsSearchItems) SetOnlineEndTime(v int64) *TaobaoJuItemsSearchItems {
    s.OnlineEndTime = &v
    return s
}
func (s *TaobaoJuItemsSearchItems) SetPicUrlForPC(v string) *TaobaoJuItemsSearchItems {
    s.PicUrlForPC = &v
    return s
}
func (s *TaobaoJuItemsSearchItems) SetPriceUspList(v []string) *TaobaoJuItemsSearchItems {
    s.PriceUspList = &v
    return s
}
func (s *TaobaoJuItemsSearchItems) SetPayPostage(v bool) *TaobaoJuItemsSearchItems {
    s.PayPostage = &v
    return s
}
