package workbench

func fixedProjects() []CourseProject {
	return []CourseProject{
		{
			ID:          "course-001",
			Title:       "连锁门店食品安全开店巡检",
			Goal:        "门店新任主管能够在30分钟内完成开店巡检，识别温控与交叉污染风险并提交整改记录",
			Institution: "启程零售培训中心",
			Chapters: []Chapter{
				{
					ID:                      "course-001-ch1",
					Title:                   "风险识别与巡检准备",
					DurationMinutes:         12,
					InstructorScript:        []string{"开店巡检不是走表，而是确认食品进入售卖环节前的最后一道防线。", "请先核对昨日整改项，再按动线准备温度计、试纸和留样标签。"},
					DemoAssets:              []DemoAsset{{Name: "开店巡检工具包摆台", Kind: "操作照片", CuePoint: "展示校准标签与消毒后的探针"}},
					ExercisePrompt:          "从工具包清单中找出两项缺失物资，并说明可能漏检的风险。",
					ExpectedExerciseMinutes: 3,
				},
				{
					ID:                      "course-001-ch2",
					Title:                   "冷藏设备温度核验",
					DurationMinutes:         10,
					InstructorScript:        []string{"先读设备显示值，再用校准探针测量高风险食品中心温度。", "两次读数不一致时，以探针复测结果启动异常处置。"},
					DemoAssets:              []DemoAsset{{Name: "冷柜探针测温示范", Kind: "操作视频", CuePoint: "探针插入最厚处并等待数值稳定"}},
					ExercisePrompt:          "冷柜显示4摄氏度、即食沙拉实测9摄氏度，请写出前三步处置动作。",
					ExpectedExerciseMinutes: 4,
				},
				{
					ID:                      "course-001-ch3",
					Title:                   "交叉污染现场处置",
					DurationMinutes:         14,
					InstructorScript:        []string{"发现生熟混放时，先隔离受影响批次，再追溯接触面和操作人员。", "处置记录必须包含批次、数量、现场照片和复核人。"},
					DemoAssets:              []DemoAsset{{Name: "冷库货架错误陈列案例", Kind: "情境图片", CuePoint: "圈出生食滴液影响范围"}},
					ExercisePrompt:          "根据货架照片标出污染路径，并口述隔离、清洁、上报的完整顺序。",
					ExpectedExerciseMinutes: 5,
				},
				{
					ID:                      "course-001-ch4",
					Title:                   "闭环复核与整改提报",
					DurationMinutes:         9,
					InstructorScript:        []string{"巡检结束前逐项确认责任人和完成时限。", "主管签字只代表复核完成，不能替代整改证据。"},
					DemoAssets:              []DemoAsset{{Name: "门店整改工单填写", Kind: "录屏", CuePoint: "补充责任人、截止时间和复核照片"}},
					ExercisePrompt:          "把一条模糊的整改描述改写为可验收的工单任务。",
					ExpectedExerciseMinutes: 3,
				},
			},
			TotalDurationMinutes: 45,
			UpdatedSequence:      1,
		},
		{
			ID:          "course-002",
			Title:       "工业设备上锁挂牌复训",
			Goal:        "维修班组长能够组织能源隔离确认，识别遗漏能源并完成上锁挂牌复核",
			Institution: "启程零售培训中心",
			Chapters: []Chapter{
				{
					ID:                      "course-002-ch1",
					Title:                   "作业前能源辨识",
					DurationMinutes:         11,
					InstructorScript:        []string{"能源清单要覆盖电力、气压、液压、重力和残余动能。"},
					DemoAssets:              []DemoAsset{{Name: "灌装线能源点位图", Kind: "设备图", CuePoint: "逐项点亮五类能源"}},
					ExercisePrompt:          "在点位图中找出被遗漏的储气罐残余压力。",
					ExpectedExerciseMinutes: 4,
				},
				{
					ID:                      "course-002-ch2",
					Title:                   "个人锁具与挂牌操作",
					DurationMinutes:         15,
					InstructorScript:        []string{"每位作业人员使用自己的锁，钥匙由本人保管。"},
					DemoAssets:              []DemoAsset{{Name: "多人锁箱操作", Kind: "操作视频", CuePoint: "班组成员依次挂个人锁"}},
					ExercisePrompt:          "三人检修中一人提前离场，判断能否代为拆锁并说明升级流程。",
					ExpectedExerciseMinutes: 5,
				},
				{
					ID:                      "course-002-ch3",
					Title:                   "零能量验证与交接",
					DurationMinutes:         12,
					InstructorScript:        []string{"隔离完成后必须尝试启动并测量确认，不能只看开关位置。"},
					DemoAssets:              []DemoAsset{{Name: "零能量验证清单", Kind: "表单演示", CuePoint: "记录验证方式和第二复核人"}},
					ExercisePrompt:          "为夜班交接补齐验证证据和签字责任。",
					ExpectedExerciseMinutes: 4,
				},
			},
			TotalDurationMinutes: 38,
			UpdatedSequence:      2,
		},
	}
}

func fixedExports() []ExportRecord {
	return []ExportRecord{
		{
			ID:        "export-001",
			ProjectID: "course-001",
			FileName:  "course-001-v1.pptx",
			Format:    "pptx",
			Status:    "已完成",
			Sequence:  3,
		},
	}
}
