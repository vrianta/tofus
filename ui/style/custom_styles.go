package style

var CardStyle = Context{
	BackgroundColor: Colors.White(),
	Padding:         EdgeInsets.SetAll(Sizes.Px(16)),
	BorderRadius: BorderRadius{
		TopLeft:     Sizes.Px(12),
		TopRight:    Sizes.Px(12),
		BottomLeft:  Sizes.Px(12),
		BottomRight: Sizes.Px(12),
	},
}
