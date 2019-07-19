// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    data, err := UnmarshalData(bytes)
//    bytes, err = data.Marshal()

package main
import "bytes"
import "errors"
import "encoding/json"
type Data []Datum


type Datum struct {
	CreatedAt            string                 `json:"created_at"`                    
	ID                   float64                `json:"id"`                            
	IDStr                string                 `json:"id_str"`                        
	Text                 string                 `json:"text"`                          
	Truncated            bool                   `json:"truncated"`                     
	Entities             DatumEntities          `json:"entities"`                      
	Source               string                 `json:"source"`                        
	InReplyToStatusID    *float64               `json:"in_reply_to_status_id"`         
	InReplyToStatusIDStr *string                `json:"in_reply_to_status_id_str"`     
	InReplyToUserID      *float64               `json:"in_reply_to_user_id"`           
	InReplyToUserIDStr   *string                `json:"in_reply_to_user_id_str"`       
	InReplyToScreenName  *string                `json:"in_reply_to_screen_name"`       
	User                 User                   `json:"user"`                          
	Geo                  interface{}            `json:"geo"`                           
	Coordinates          interface{}            `json:"coordinates"`                   
	Place                interface{}            `json:"place"`                         
	Contributors         interface{}            `json:"contributors"`                  
	RetweetedStatus      *RetweetedStatus       `json:"retweeted_status,omitempty"`    
	IsQuoteStatus        bool                   `json:"is_quote_status"`               
	RetweetCount         int64                  `json:"retweet_count"`                 
	FavoriteCount        int64                  `json:"favorite_count"`                
	Favorited            bool                   `json:"favorited"`                     
	Retweeted            bool                   `json:"retweeted"`                     
	Lang                 Lang                   `json:"lang"`                          
	ExtendedEntities     *DatumExtendedEntities `json:"extended_entities,omitempty"`   
	PossiblySensitive    *bool                  `json:"possibly_sensitive,omitempty"`  
	QuotedStatusID       *float64               `json:"quoted_status_id,omitempty"`    
	QuotedStatusIDStr    *string                `json:"quoted_status_id_str,omitempty"`
	QuotedStatus         *QuotedStatus          `json:"quoted_status,omitempty"`       
}

type DatumEntities struct {
	Hashtags     []Hashtag     `json:"hashtags"`     
	Symbols      []interface{} `json:"symbols"`      
	UserMentions []UserMention `json:"user_mentions"`
	Urls         []URL         `json:"urls"`         
	Media        []PurpleMedia `json:"media"`        
}

type Hashtag struct {
	Text    string  `json:"text"`   
	Indices []int64 `json:"indices"`
}

type PurpleMedia struct {
	ID                float64  `json:"id"`                            
	IDStr             string   `json:"id_str"`                        
	Indices           []int64  `json:"indices"`                       
	MediaURL          string   `json:"media_url"`                     
	MediaURLHTTPS     string   `json:"media_url_https"`               
	URL               string   `json:"url"`                           
	DisplayURL        string   `json:"display_url"`                   
	ExpandedURL       string   `json:"expanded_url"`                  
	Type              string   `json:"type"`                          
	Sizes             Sizes    `json:"sizes"`                         
	SourceStatusID    *float64 `json:"source_status_id,omitempty"`    
	SourceStatusIDStr *string  `json:"source_status_id_str,omitempty"`
	SourceUserID      *int64   `json:"source_user_id,omitempty"`      
	SourceUserIDStr   *string  `json:"source_user_id_str,omitempty"`  
}

type Sizes struct {
	Thumb  Large `json:"thumb"` 
	Medium Large `json:"medium"`
	Large  Large `json:"large"` 
	Small  Large `json:"small"` 
}

type Large struct {
	W      int64  `json:"w"`     
	H      int64  `json:"h"`     
	Resize Resize `json:"resize"`
}

type URL struct {
	URL         string  `json:"url"`         
	ExpandedURL string  `json:"expanded_url"`
	DisplayURL  string  `json:"display_url"` 
	Indices     []int64 `json:"indices"`     
}

type UserMention struct {
	ScreenName string  `json:"screen_name"`
	Name       string  `json:"name"`       
	ID         float64 `json:"id"`         
	IDStr      string  `json:"id_str"`     
	Indices    []int64 `json:"indices"`    
}

type DatumExtendedEntities struct {
	Media []PurpleMedia `json:"media"`
}

type QuotedStatus struct {
	CreatedAt            string                       `json:"created_at"`               
	ID                   float64                      `json:"id"`                       
	IDStr                string                       `json:"id_str"`                   
	Text                 string                       `json:"text"`                     
	Truncated            bool                         `json:"truncated"`                
	Entities             QuotedStatusEntities         `json:"entities"`                 
	ExtendedEntities     QuotedStatusExtendedEntities `json:"extended_entities"`        
	Source               string                       `json:"source"`                   
	InReplyToStatusID    interface{}                  `json:"in_reply_to_status_id"`    
	InReplyToStatusIDStr interface{}                  `json:"in_reply_to_status_id_str"`
	InReplyToUserID      interface{}                  `json:"in_reply_to_user_id"`      
	InReplyToUserIDStr   interface{}                  `json:"in_reply_to_user_id_str"`  
	InReplyToScreenName  interface{}                  `json:"in_reply_to_screen_name"`  
	User                 User                         `json:"user"`                     
	Geo                  interface{}                  `json:"geo"`                      
	Coordinates          interface{}                  `json:"coordinates"`              
	Place                interface{}                  `json:"place"`                    
	Contributors         interface{}                  `json:"contributors"`             
	IsQuoteStatus        bool                         `json:"is_quote_status"`          
	RetweetCount         int64                        `json:"retweet_count"`            
	FavoriteCount        int64                        `json:"favorite_count"`           
	Favorited            bool                         `json:"favorited"`                
	Retweeted            bool                         `json:"retweeted"`                
	PossiblySensitive    bool                         `json:"possibly_sensitive"`       
	Lang                 Lang                         `json:"lang"`                     
}

type QuotedStatusEntities struct {
	Hashtags     []Hashtag     `json:"hashtags"`     
	Symbols      []interface{} `json:"symbols"`      
	UserMentions []UserMention `json:"user_mentions"`
	Urls         []URL         `json:"urls"`         
	Media        []FluffyMedia `json:"media"`        
}

type FluffyMedia struct {
	ID                  float64              `json:"id"`                             
	IDStr               string               `json:"id_str"`                         
	Indices             []int64              `json:"indices"`                        
	MediaURL            string               `json:"media_url"`                      
	MediaURLHTTPS       string               `json:"media_url_https"`                
	URL                 string               `json:"url"`                            
	DisplayURL          string               `json:"display_url"`                    
	ExpandedURL         string               `json:"expanded_url"`                   
	Type                string               `json:"type"`                           
	Sizes               Sizes                `json:"sizes"`                          
	VideoInfo           *VideoInfo           `json:"video_info,omitempty"`           
	AdditionalMediaInfo *AdditionalMediaInfo `json:"additional_media_info,omitempty"`
}

type AdditionalMediaInfo struct {
	Monetizable bool `json:"monetizable"`
}

type VideoInfo struct {
	AspectRatio    []int64   `json:"aspect_ratio"`   
	DurationMillis int64     `json:"duration_millis"`
	Variants       []Variant `json:"variants"`       
}

type Variant struct {
	ContentType string `json:"content_type"`     
	URL         string `json:"url"`              
	Bitrate     *int64 `json:"bitrate,omitempty"`
}

type QuotedStatusExtendedEntities struct {
	Media []FluffyMedia `json:"media"`
}

type User struct {
	ID                             *ID                       `json:"id"`                                
	IDStr                          string                    `json:"id_str"`                            
	Name                           string                    `json:"name"`                              
	ScreenName                     string                    `json:"screen_name"`                       
	Location                       string                    `json:"location"`                          
	Description                    string                    `json:"description"`                       
	URL                            *string                   `json:"url"`                               
	Entities                       UserEntities              `json:"entities"`                          
	Protected                      bool                      `json:"protected"`                         
	FollowersCount                 int64                     `json:"followers_count"`                   
	FriendsCount                   int64                     `json:"friends_count"`                     
	ListedCount                    int64                     `json:"listed_count"`                      
	CreatedAt                      string                    `json:"created_at"`                        
	FavouritesCount                int64                     `json:"favourites_count"`                  
	UTCOffset                      interface{}               `json:"utc_offset"`                        
	TimeZone                       interface{}               `json:"time_zone"`                         
	GeoEnabled                     bool                      `json:"geo_enabled"`                       
	Verified                       bool                      `json:"verified"`                          
	StatusesCount                  int64                     `json:"statuses_count"`                    
	Lang                           interface{}               `json:"lang"`                              
	ContributorsEnabled            bool                      `json:"contributors_enabled"`              
	IsTranslator                   bool                      `json:"is_translator"`                     
	IsTranslationEnabled           bool                      `json:"is_translation_enabled"`            
	ProfileBackgroundColor         ProfileBackgroundColor    `json:"profile_background_color"`          
	ProfileBackgroundImageURL      string                    `json:"profile_background_image_url"`      
	ProfileBackgroundImageURLHTTPS string                    `json:"profile_background_image_url_https"`
	ProfileBackgroundTile          bool                      `json:"profile_background_tile"`           
	ProfileImageURL                string                    `json:"profile_image_url"`                 
	ProfileImageURLHTTPS           string                    `json:"profile_image_url_https"`           
	ProfileBannerURL               string                    `json:"profile_banner_url"`                
	ProfileLinkColor               string                    `json:"profile_link_color"`                
	ProfileSidebarBorderColor      ProfileSidebarBorderColor `json:"profile_sidebar_border_color"`      
	ProfileSidebarFillColor        ProfileSidebarFillColor   `json:"profile_sidebar_fill_color"`        
	ProfileTextColor               string                    `json:"profile_text_color"`                
	ProfileUseBackgroundImage      bool                      `json:"profile_use_background_image"`      
	HasExtendedProfile             bool                      `json:"has_extended_profile"`              
	DefaultProfile                 bool                      `json:"default_profile"`                   
	DefaultProfileImage            bool                      `json:"default_profile_image"`             
	Following                      interface{}               `json:"following"`                         
	FollowRequestSent              interface{}               `json:"follow_request_sent"`               
	Notifications                  interface{}               `json:"notifications"`                     
	TranslatorType                 TranslatorType            `json:"translator_type"`                   
}

type UserEntities struct {
	URL         *Description `json:"url,omitempty"`
	Description Description  `json:"description"`  
}

type Description struct {
	Urls []URL `json:"urls"`
}

type RetweetedStatus struct {
	CreatedAt            string                        `json:"created_at"`                  
	ID                   float64                       `json:"id"`                          
	IDStr                string                        `json:"id_str"`                      
	Text                 string                        `json:"text"`                        
	Truncated            bool                          `json:"truncated"`                   
	Entities             QuotedStatusEntities          `json:"entities"`                    
	Source               string                        `json:"source"`                      
	InReplyToStatusID    interface{}                   `json:"in_reply_to_status_id"`       
	InReplyToStatusIDStr interface{}                   `json:"in_reply_to_status_id_str"`   
	InReplyToUserID      interface{}                   `json:"in_reply_to_user_id"`         
	InReplyToUserIDStr   interface{}                   `json:"in_reply_to_user_id_str"`     
	InReplyToScreenName  interface{}                   `json:"in_reply_to_screen_name"`     
	User                 User                          `json:"user"`                        
	Geo                  interface{}                   `json:"geo"`                         
	Coordinates          interface{}                   `json:"coordinates"`                 
	Place                interface{}                   `json:"place"`                       
	Contributors         interface{}                   `json:"contributors"`                
	IsQuoteStatus        bool                          `json:"is_quote_status"`             
	RetweetCount         int64                         `json:"retweet_count"`               
	FavoriteCount        int64                         `json:"favorite_count"`              
	Favorited            bool                          `json:"favorited"`                   
	Retweeted            bool                          `json:"retweeted"`                   
	PossiblySensitive    *bool                         `json:"possibly_sensitive,omitempty"`
	Lang                 Lang                          `json:"lang"`                        
	ExtendedEntities     *QuotedStatusExtendedEntities `json:"extended_entities,omitempty"` 
}

type Resize string
const (
	Crop Resize = "crop"
	Fit Resize = "fit"
)

type Lang string
const (
	En Lang = "en"
	HT Lang = "ht"
	Und Lang = "und"
)

type ProfileBackgroundColor string
const (
	Ebebeb ProfileBackgroundColor = "EBEBEB"
	Edece9 ProfileBackgroundColor = "EDECE9"
	ProfileBackgroundColor000000 ProfileBackgroundColor = "000000"
	ProfileBackgroundColorC0DEED ProfileBackgroundColor = "C0DEED"
)

type ProfileSidebarBorderColor string
const (
	D3D2CF ProfileSidebarBorderColor = "D3D2CF"
	Ffffff ProfileSidebarBorderColor = "FFFFFF"
	ProfileSidebarBorderColor000000 ProfileSidebarBorderColor = "000000"
	ProfileSidebarBorderColorC0DEED ProfileSidebarBorderColor = "C0DEED"
)

type ProfileSidebarFillColor string
const (
	Ddeef6 ProfileSidebarFillColor = "DDEEF6"
	E3E2De ProfileSidebarFillColor = "E3E2DE"
	ProfileSidebarFillColor000000 ProfileSidebarFillColor = "000000"
)

type TranslatorType string
const (
	None TranslatorType = "none"
	Regular TranslatorType = "regular"
)

type ID struct {
	Double  *float64
	Integer *int64
}
func (x *ID) UnmarshalJSON(data []byte) error {
	object, err := unmarshalUnion(data, &x.Integer, &x.Double, nil, nil, false, nil, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *ID) MarshalJSON() ([]byte, error) {
	return marshalUnion(x.Integer, x.Double, nil, nil, false, nil, false, nil, false, nil, false, nil, false)
}

func unmarshalUnion(data []byte, pi **int64, pf **float64, pb **bool, ps **string, haveArray bool, pa interface{}, haveObject bool, pc interface{}, haveMap bool, pm interface{}, haveEnum bool, pe interface{}, nullable bool) (bool, error) {
	if pi != nil {
		*pi = nil
	}
	if pf != nil {
		*pf = nil
	}
	if pb != nil {
		*pb = nil
	}
	if ps != nil {
		*ps = nil
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()
	tok, err := dec.Token()
	if err != nil {
		return false, err
	}

	switch v := tok.(type) {
	case json.Number:
		if pi != nil {
			i, err := v.Int64()
			if err == nil {
				*pi = &i
				return false, nil
			}
		}
		if pf != nil {
			f, err := v.Float64()
			if err == nil {
				*pf = &f
				return false, nil
			}
			return false, errors.New("Unparsable number")
		}
		return false, errors.New("Union does not contain number")
	case float64:
		return false, errors.New("Decoder should not return float64")
	case bool:
		if pb != nil {
			*pb = &v
			return false, nil
		}
		return false, errors.New("Union does not contain bool")
	case string:
		if haveEnum {
			return false, json.Unmarshal(data, pe)
		}
		if ps != nil {
			*ps = &v
			return false, nil
		}
		return false, errors.New("Union does not contain string")
	case nil:
		if nullable {
			return false, nil
		}
		return false, errors.New("Union does not contain null")
	case json.Delim:
		if v == '{' {
			if haveObject {
				return true, json.Unmarshal(data, pc)
			}
			if haveMap {
				return false, json.Unmarshal(data, pm)
			}
			return false, errors.New("Union does not contain object")
		}
		if v == '[' {
			if haveArray {
				return false, json.Unmarshal(data, pa)
			}
			return false, errors.New("Union does not contain array")
		}
		return false, errors.New("Cannot handle delimiter")
	}
	return false, errors.New("Cannot unmarshal union")

}

func marshalUnion(pi *int64, pf *float64, pb *bool, ps *string, haveArray bool, pa interface{}, haveObject bool, pc interface{}, haveMap bool, pm interface{}, haveEnum bool, pe interface{}, nullable bool) ([]byte, error) {
	if pi != nil {
		return json.Marshal(*pi)
	}
	if pf != nil {
		return json.Marshal(*pf)
	}
	if pb != nil {
		return json.Marshal(*pb)
	}
	if ps != nil {
		return json.Marshal(*ps)
	}
	if haveArray {
		return json.Marshal(pa)
	}
	if haveObject {
		return json.Marshal(pc)
	}
	if haveMap {
		return json.Marshal(pm)
	}
	if haveEnum {
		return json.Marshal(pe)
	}
	if nullable {
		return json.Marshal(nil)
	}
	return nil, errors.New("Union must not be null")
}

