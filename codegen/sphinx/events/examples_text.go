package main

var eventExampleStr = map[string]string{
	"server.panic": `{
	"event": "server.panic",
	"params": {
		"code": "INVALID_TEAM",
		"debug": null
	}
}`,
	"server.uisettings": `{
	"last-tab": 91238475
}`,
	"server.warning": `{
	"event": "server.warning",
	"params": {
		"message": "unknown event: client.deeeeeaddddbeeeeeef",
		"orig": {
			"confirm_id": "c6280a82ed1c",
			"event": "client.deeeeeaddddbeeeeeef"
		}
	}
}`,
	"server.confirm": `{
	"event": "server.confirm",
	"params": {
		"confirm_id": "b8b2ccd6-35a6-408f-a591-c696a9f9e73e"
	}
}`,
	"server.roster": `{
    "event": "server.roster",
    "params": {
        "badge": 449,
        "contact_sections": [
            {
                "gentime": 1568764805859444992,
                "name": "Боты",
                "sort_ordering": 99,
                "uid": "3a8b23f2-25b4-4acb-afab-590303bedd50"
            }
        ],
        "contacts": [
            {
                "botname": "systembot",
                "can_send_message": true,
                "changeable_fields": [],
                "contact_email": "support@tada.team",
                "contact_phone": "",
                "display_name": "tada.teamBot",
                "icons": {
                    "lg": {
                        "height": 512,
                        "url": "http://127.0.0.1:8000/static/tada-bots/systembot512.png",
                        "width": 512
                    },
                    "sm": {
                        "height": 256,
                        "url": "http://127.0.0.1:8000/static/tada-bots/systembot256.png",
                        "width": 256
                    }
                },
                "jid": "d-23eae5d4-4371-4a43-8815-a6a655cb623e",
                "last_activity": "2019-09-27T01:01:32.354349Z",
                "role": "Виртуальный ассистент",
                "section": "3a8b23f2-25b4-4acb-afab-590303bedd50",
                "sections": [
                    "3a8b23f2-25b4-4acb-afab-590303bedd50"
                ],
                "status": "member"
            },
            {
                "can_add_to_group": true,
                "can_call": true,
                "can_create_task": true,
                "can_send_message": true,
                "changeable_fields": [],
                "contact_email": "",
                "contact_phone": "+75550000000",
                "display_name": "Контакт №1",
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/578e3c/%D0%9A%E2%84%96/256.png"
                },
                "jid": "d-181b9859-ea3f-4e26-85c9-9e7ee7aaaa9a",
                "last_activity": null,
                "role": "",
                "sections": [],
                "status": "admin"
            },
            {
                "can_add_to_group": true,
                "can_call": true,
                "can_create_task": true,
                "can_send_message": true,
                "changeable_fields": [],
                "contact_email": "",
                "contact_phone": "+75550000009",
                "display_name": "Контакт №10",
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/579dda/10/256.png"
                },
                "jid": "d-a87e168a-78df-49aa-9abb-9b5409eb7de5",
                "last_activity": null,
                "role": "",
                "sections": [],
                "status": "member"
            },
            {
                "can_add_to_group": true,
                "can_call": true,
                "can_create_task": true,
                "can_send_message": true,
                "changeable_fields": [],
                "contact_email": "",
                "contact_phone": "+75550000010",
                "display_name": "Контакт №11",
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/8e88eb/11/256.png"
                },
                "jid": "d-8bb2ed5e-5cf4-457c-9a7b-070d7e4116a9",
                "last_activity": null,
                "role": "",
                "sections": [],
                "status": "member"
            },
            {
                "can_add_to_group": true,
                "can_call": true,
                "can_create_task": true,
                "can_send_message": true,
                "changeable_fields": [],
                "contact_email": "",
                "contact_phone": "+75550000011",
                "display_name": "Контакт №12",
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/e36659/12/256.png"
                },
                "jid": "d-5049173a-93dc-4904-8796-6807b6f07bf1",
                "last_activity": null,
                "role": "",
                "sections": [],
                "status": "member"
            },
            {
                "can_add_to_group": true,
                "can_call": true,
                "can_create_task": true,
                "can_send_message": true,
                "changeable_fields": [],
                "contact_email": "",
                "contact_phone": "+75550000012",
                "display_name": "Контакт №13",
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/f0769a/13/256.png"
                },
                "jid": "d-26879ae0-a4b9-4523-8875-4e26c4787d27",
                "last_activity": null,
                "role": "",
                "sections": [],
                "status": "member"
            },
            {
                "can_add_to_group": true,
                "can_call": true,
                "can_create_task": true,
                "can_send_message": true,
                "changeable_fields": [],
                "contact_email": "",
                "contact_phone": "+75550000013",
                "display_name": "Контакт №14",
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/f08c50/14/256.png"
                },
                "jid": "d-84e3cb48-598b-4ffd-918e-0662ebeb3469",
                "last_activity": null,
                "role": "",
                "sections": [],
                "status": "member"
            },
            {
                "can_add_to_group": true,
                "can_call": true,
                "can_create_task": true,
                "can_send_message": true,
                "changeable_fields": [],
                "contact_email": "",
                "contact_phone": "+75550000014",
                "display_name": "Контакт №15",
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/578e3c/15/256.png"
                },
                "jid": "d-6496b27c-3251-4683-896a-a7f2191b7a33",
                "last_activity": null,
                "role": "",
                "sections": [],
                "status": "member"
            },
            {
                "can_add_to_group": true,
                "can_call": true,
                "can_create_task": true,
                "can_send_message": true,
                "changeable_fields": [],
                "contact_email": "",
                "contact_phone": "+75550000015",
                "display_name": "Контакт №16",
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/63bed3/16/256.png"
                },
                "jid": "d-c540dd6e-4da3-4b65-8e89-bd35554978c2",
                "last_activity": null,
                "role": "",
                "sections": [],
                "status": "member"
            },
            {
                "can_add_to_group": true,
                "can_call": true,
                "can_create_task": true,
                "can_send_message": true,
                "changeable_fields": [],
                "contact_email": "",
                "contact_phone": "+75550000016",
                "display_name": "Контакт №17",
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/579dda/17/256.png"
                },
                "jid": "d-951775af-15fd-4e5f-bb13-099c9afd6586",
                "last_activity": null,
                "role": "",
                "sections": [],
                "status": "member"
            },
            {
                "can_add_to_group": true,
                "can_call": true,
                "can_create_task": true,
                "can_send_message": true,
                "changeable_fields": [],
                "contact_email": "",
                "contact_phone": "+75550000017",
                "display_name": "Контакт №18",
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/8e88eb/18/256.png"
                },
                "jid": "d-ce8117ae-867d-400d-b7c8-e7df6a2ecee7",
                "last_activity": null,
                "role": "",
                "sections": [],
                "status": "member"
            },
            {
                "can_add_to_group": true,
                "can_call": true,
                "can_create_task": true,
                "can_send_message": true,
                "changeable_fields": [],
                "contact_email": "",
                "contact_phone": "+75550000018",
                "display_name": "Контакт №19",
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/e36659/19/256.png"
                },
                "jid": "d-1d60b554-5da7-4f17-a078-eb77168081a8",
                "last_activity": null,
                "role": "",
                "sections": [],
                "status": "member"
            },
            {
                "alt_send": false,
                "can_add_to_group": true,
                "can_create_task": true,
                "can_delete": true,
                "changeable_fields": [
                    "alt_send",
                    "contact_email",
                    "contact_mshort_view",
                    "contact_phone",
                    "contact_short_view",
                    "contact_show_archived",
                    "debug_show_activity",
                    "default_lang",
                    "family_name",
                    "given_name",
                    "group_mshort_view",
                    "group_notifications_enabled",
                    "group_short_view",
                    "icons",
                    "phone",
                    "role",
                    "task_mshort_view",
                    "task_notifications_enabled",
                    "task_short_view",
                    "unread_first"
                ],
                "contact_email": "",
                "contact_mshort_view": false,
                "contact_phone": "+75550000001",
                "contact_short_view": false,
                "contact_show_archived": false,
                "debug_show_activity": false,
                "default_lang": null,
                "display_name": "Контакт №2",
                "family_name": "№2",
                "given_name": "Контакт",
                "group_mshort_view": false,
                "group_notifications_enabled": true,
                "group_short_view": false,
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/63bed3/%D0%9A%E2%84%96/256.png"
                },
                "jid": "d-bd500a50-3a38-44d1-a43c-fb1a48e1a79e",
                "last_activity": null,
                "role": "",
                "sections": [],
                "status": "member",
                "task_mshort_view": false,
                "task_notifications_enabled": true,
                "task_short_view": false,
                "unread_first": false
            },
            {
                "can_add_to_group": true,
                "can_call": true,
                "can_create_task": true,
                "can_send_message": true,
                "changeable_fields": [],
                "contact_email": "",
                "contact_phone": "+75550000019",
                "display_name": "Контакт №20",
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/f0769a/20/256.png"
                },
                "jid": "d-13b47a4f-4473-4a8f-82ef-00b00ce5c02f",
                "last_activity": null,
                "role": "",
                "sections": [],
                "status": "member"
            },
            {
                "can_add_to_group": true,
                "can_call": true,
                "can_create_task": true,
                "can_send_message": true,
                "changeable_fields": [],
                "contact_email": "",
                "contact_phone": "+75550000020",
                "display_name": "Контакт №21",
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/f08c50/21/256.png"
                },
                "jid": "d-14ba8969-e3a9-4daf-b6c3-bed4ea931e06",
                "last_activity": null,
                "role": "",
                "sections": [],
                "status": "member"
            },
            {
                "can_add_to_group": true,
                "can_call": true,
                "can_create_task": true,
                "can_send_message": true,
                "changeable_fields": [],
                "contact_email": "",
                "contact_phone": "+75550000021",
                "display_name": "Контакт №22",
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/578e3c/22/256.png"
                },
                "jid": "d-571d1843-96d1-497b-a53c-2cc03d815194",
                "last_activity": null,
                "role": "",
                "sections": [],
                "status": "member"
            },
            {
                "can_add_to_group": true,
                "can_call": true,
                "can_create_task": true,
                "can_send_message": true,
                "changeable_fields": [],
                "contact_email": "",
                "contact_phone": "+75550000022",
                "display_name": "Контакт №23",
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/63bed3/23/256.png"
                },
                "jid": "d-c15b9ad4-b03b-4dc1-aeaa-725e09bedbcf",
                "last_activity": null,
                "role": "",
                "sections": [],
                "status": "member"
            },
            {
                "can_add_to_group": true,
                "can_call": true,
                "can_create_task": true,
                "can_send_message": true,
                "changeable_fields": [],
                "contact_email": "",
                "contact_phone": "+75550000023",
                "display_name": "Контакт №24",
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/579dda/24/256.png"
                },
                "jid": "d-0f902404-96c4-4c5e-8385-d33a46fd5566",
                "last_activity": null,
                "role": "",
                "sections": [],
                "status": "member"
            },
            {
                "can_add_to_group": true,
                "can_call": true,
                "can_create_task": true,
                "can_send_message": true,
                "changeable_fields": [],
                "contact_email": "",
                "contact_phone": "+75550000024",
                "display_name": "Контакт №25",
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/8e88eb/25/256.png"
                },
                "jid": "d-b54d8559-4d94-46e6-b1ff-673a55e5cf5d",
                "last_activity": null,
                "role": "",
                "sections": [],
                "status": "member"
            },
            {
                "can_add_to_group": true,
                "can_call": true,
                "can_create_task": true,
                "can_send_message": true,
                "changeable_fields": [],
                "contact_email": "",
                "contact_phone": "+75550000025",
                "display_name": "Контакт №26",
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/e36659/26/256.png"
                },
                "jid": "d-d62ca7c6-83b5-4853-bfc5-10e8c2e68b13",
                "last_activity": null,
                "role": "",
                "sections": [],
                "status": "member"
            },
            {
                "can_add_to_group": true,
                "can_call": true,
                "can_create_task": true,
                "can_send_message": true,
                "changeable_fields": [],
                "contact_email": "",
                "contact_phone": "+75550000026",
                "display_name": "Контакт №27",
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/f0769a/27/256.png"
                },
                "jid": "d-7589ac63-4c0d-4685-9bf5-63b83cdcc3ab",
                "last_activity": null,
                "role": "",
                "sections": [],
                "status": "member"
            },
            {
                "can_add_to_group": true,
                "can_call": true,
                "can_create_task": true,
                "can_send_message": true,
                "changeable_fields": [],
                "contact_email": "",
                "contact_phone": "+75550000027",
                "display_name": "Контакт №28",
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/f08c50/28/256.png"
                },
                "jid": "d-10e10caf-6127-4649-b2d6-5ad1395d32a5",
                "last_activity": null,
                "role": "",
                "sections": [],
                "status": "member"
            },
            {
                "can_add_to_group": true,
                "can_call": true,
                "can_create_task": true,
                "can_send_message": true,
                "changeable_fields": [],
                "contact_email": "",
                "contact_phone": "+75550000028",
                "display_name": "Контакт №29",
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/578e3c/29/256.png"
                },
                "jid": "d-1b20a9aa-ecf5-49fd-8477-0998d42e1ade",
                "last_activity": null,
                "role": "",
                "sections": [],
                "status": "member"
            },
            {
                "can_add_to_group": true,
                "can_call": true,
                "can_create_task": true,
                "can_send_message": true,
                "changeable_fields": [],
                "contact_email": "",
                "contact_phone": "+75550000002",
                "display_name": "Контакт №3",
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/579dda/%D0%9A%E2%84%96/256.png"
                },
                "jid": "d-ee3d7c07-fe0f-4cbd-bf3a-b8e72f0ff1c6",
                "last_activity": null,
                "role": "",
                "sections": [],
                "status": "member"
            },
            {
                "can_add_to_group": true,
                "can_call": true,
                "can_create_task": true,
                "can_send_message": true,
                "changeable_fields": [],
                "contact_email": "",
                "contact_phone": "+75550000029",
                "display_name": "Контакт №30",
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/63bed3/30/256.png"
                },
                "jid": "d-bc7550f7-ae4b-41de-82d3-d4b5b5b7f600",
                "last_activity": null,
                "role": "",
                "sections": [],
                "status": "member"
            },
            {
                "can_add_to_group": true,
                "can_call": true,
                "can_create_task": true,
                "can_send_message": true,
                "changeable_fields": [],
                "contact_email": "",
                "contact_phone": "+75550000030",
                "display_name": "Контакт №31",
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/579dda/31/256.png"
                },
                "jid": "d-bfe1fa23-c32b-48ea-8fed-d2b6f583fb77",
                "last_activity": null,
                "role": "",
                "sections": [],
                "status": "member"
            },
            {
                "can_add_to_group": true,
                "can_call": true,
                "can_create_task": true,
                "can_send_message": true,
                "changeable_fields": [],
                "contact_email": "",
                "contact_phone": "+75550000031",
                "display_name": "Контакт №32",
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/8e88eb/32/256.png"
                },
                "jid": "d-a792832a-d9f6-4bbc-b545-f2800e329c8c",
                "last_activity": null,
                "role": "",
                "sections": [],
                "status": "member"
            },
            {
                "can_add_to_group": true,
                "can_call": true,
                "can_create_task": true,
                "can_send_message": true,
                "changeable_fields": [],
                "contact_email": "",
                "contact_phone": "+75550000003",
                "display_name": "Контакт №4",
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/8e88eb/%D0%9A%E2%84%96/256.png"
                },
                "jid": "d-c094f142-1b54-4fae-8128-9b41504b56a9",
                "last_activity": null,
                "role": "",
                "sections": [],
                "status": "member"
            },
            {
                "can_add_to_group": true,
                "can_call": true,
                "can_create_task": true,
                "can_send_message": true,
                "changeable_fields": [],
                "contact_email": "",
                "contact_phone": "+75550000004",
                "display_name": "Контакт №5",
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/e36659/%D0%9A%E2%84%96/256.png"
                },
                "jid": "d-53e898d1-3d7b-4cc4-9ff6-ca034b5709cb",
                "last_activity": null,
                "role": "",
                "sections": [],
                "status": "member"
            },
            {
                "can_add_to_group": true,
                "can_call": true,
                "can_create_task": true,
                "can_send_message": true,
                "changeable_fields": [],
                "contact_email": "",
                "contact_phone": "+75550000005",
                "display_name": "Контакт №6",
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/f0769a/%D0%9A%E2%84%96/256.png"
                },
                "jid": "d-af8d0ae3-ff07-4c49-94d0-7d531f972b0e",
                "last_activity": null,
                "role": "",
                "sections": [],
                "status": "member"
            },
            {
                "can_add_to_group": true,
                "can_call": true,
                "can_create_task": true,
                "can_send_message": true,
                "changeable_fields": [],
                "contact_email": "",
                "contact_phone": "+75550000006",
                "display_name": "Контакт №7",
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/f08c50/%D0%9A%E2%84%96/256.png"
                },
                "jid": "d-f73f2847-4fcb-4f7e-bc11-8d2aaf21cf00",
                "last_activity": null,
                "role": "",
                "sections": [],
                "status": "member"
            },
            {
                "can_add_to_group": true,
                "can_call": true,
                "can_create_task": true,
                "can_send_message": true,
                "changeable_fields": [],
                "contact_email": "",
                "contact_phone": "+75550000007",
                "display_name": "Контакт №8",
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/578e3c/%D0%9A%E2%84%96/256.png"
                },
                "jid": "d-2bde58fa-3b2b-4c10-8041-b0a65e6d74b4",
                "last_activity": null,
                "role": "",
                "sections": [],
                "status": "member"
            },
            {
                "can_add_to_group": true,
                "can_call": true,
                "can_create_task": true,
                "can_send_message": true,
                "changeable_fields": [],
                "contact_email": "",
                "contact_phone": "+75550000008",
                "display_name": "Контакт №9",
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/63bed3/%D0%9A%E2%84%96/256.png"
                },
                "jid": "d-425ef763-2162-4dcf-8e0e-5ce0615518d8",
                "last_activity": null,
                "role": "",
                "sections": [],
                "status": "member"
            },
            {
                "botname": "notebot",
                "can_send_message": true,
                "changeable_fields": [],
                "contact_email": "",
                "contact_phone": "",
                "display_name": "Мои заметки",
                "icons": {
                    "lg": {
                        "height": 512,
                        "url": "http://127.0.0.1:8000/static/tada-bots/notebot512.png",
                        "width": 512
                    },
                    "sm": {
                        "height": 256,
                        "url": "http://127.0.0.1:8000/static/tada-bots/notebot256.png",
                        "width": 256
                    }
                },
                "jid": "d-6d02dd61-8c07-4d11-a1f5-4584569d7b6d",
                "last_activity": "2019-09-27T01:01:32.354092Z",
                "role": "Блокнот",
                "section": "3a8b23f2-25b4-4acb-afab-590303bedd50",
                "sections": [
                    "3a8b23f2-25b4-4acb-afab-590303bedd50"
                ],
                "status": "member"
            }
        ],
        "devices": [],
        "direct_chats": [
            {
                "can_call": true,
                "can_send_message": true,
                "chat_type": "direct",
                "counters_enabled": true,
                "created": "2019-09-18T00:00:42.318781Z",
                "display_name": "Контакт №32",
                "gentime": 1568764842326248192,
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/8e88eb/32/256.png"
                },
                "jid": "d-a792832a-d9f6-4bbc-b545-f2800e329c8c",
                "notifications_enabled": true
            },
            {
                "can_call": true,
                "can_send_message": true,
                "chat_type": "direct",
                "counters_enabled": true,
                "created": "2019-09-18T00:00:41.097659Z",
                "display_name": "Контакт №31",
                "gentime": 1568764841104200192,
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/579dda/31/256.png"
                },
                "jid": "d-bfe1fa23-c32b-48ea-8fed-d2b6f583fb77",
                "notifications_enabled": true
            },
            {
                "can_call": true,
                "can_send_message": true,
                "chat_type": "direct",
                "counters_enabled": true,
                "created": "2019-09-18T00:00:38.428242Z",
                "display_name": "Контакт №30",
                "gentime": 1568764838432389632,
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/63bed3/30/256.png"
                },
                "jid": "d-bc7550f7-ae4b-41de-82d3-d4b5b5b7f600",
                "notifications_enabled": true
            },
            {
                "can_call": true,
                "can_send_message": true,
                "chat_type": "direct",
                "counters_enabled": true,
                "created": "2019-09-18T00:00:36.676433Z",
                "display_name": "Контакт №29",
                "gentime": 1568764836680841984,
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/578e3c/29/256.png"
                },
                "jid": "d-1b20a9aa-ecf5-49fd-8477-0998d42e1ade",
                "notifications_enabled": true
            },
            {
                "can_call": true,
                "can_send_message": true,
                "chat_type": "direct",
                "counters_enabled": true,
                "created": "2019-09-18T00:00:34.945690Z",
                "display_name": "Контакт №28",
                "gentime": 1568764834949238528,
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/f08c50/28/256.png"
                },
                "jid": "d-10e10caf-6127-4649-b2d6-5ad1395d32a5",
                "notifications_enabled": true
            },
            {
                "can_call": true,
                "can_send_message": true,
                "chat_type": "direct",
                "counters_enabled": true,
                "created": "2019-09-18T00:00:33.326639Z",
                "display_name": "Контакт №27",
                "gentime": 1568764833332984320,
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/f0769a/27/256.png"
                },
                "jid": "d-7589ac63-4c0d-4685-9bf5-63b83cdcc3ab",
                "notifications_enabled": true
            },
            {
                "can_call": true,
                "can_send_message": true,
                "chat_type": "direct",
                "counters_enabled": true,
                "created": "2019-09-18T00:00:31.700784Z",
                "display_name": "Контакт №26",
                "gentime": 1568764831705233664,
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/e36659/26/256.png"
                },
                "jid": "d-d62ca7c6-83b5-4853-bfc5-10e8c2e68b13",
                "notifications_enabled": true
            },
            {
                "can_call": true,
                "can_send_message": true,
                "chat_type": "direct",
                "counters_enabled": true,
                "created": "2019-09-18T00:00:30.429840Z",
                "display_name": "Контакт №25",
                "gentime": 1568764830435267328,
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/8e88eb/25/256.png"
                },
                "jid": "d-b54d8559-4d94-46e6-b1ff-673a55e5cf5d",
                "notifications_enabled": true
            },
            {
                "can_call": true,
                "can_send_message": true,
                "chat_type": "direct",
                "counters_enabled": true,
                "created": "2019-09-18T00:00:28.982264Z",
                "display_name": "Контакт №24",
                "gentime": 1568764828988044800,
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/579dda/24/256.png"
                },
                "jid": "d-0f902404-96c4-4c5e-8385-d33a46fd5566",
                "notifications_enabled": true
            },
            {
                "can_call": true,
                "can_send_message": true,
                "chat_type": "direct",
                "counters_enabled": true,
                "created": "2019-09-18T00:00:27.429493Z",
                "display_name": "Контакт №23",
                "gentime": 1568764827437857536,
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/63bed3/23/256.png"
                },
                "jid": "d-c15b9ad4-b03b-4dc1-aeaa-725e09bedbcf",
                "notifications_enabled": true
            },
            {
                "can_call": true,
                "can_send_message": true,
                "chat_type": "direct",
                "counters_enabled": true,
                "created": "2019-09-18T00:00:26.171104Z",
                "display_name": "Контакт №22",
                "gentime": 1568764826174592768,
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/578e3c/22/256.png"
                },
                "jid": "d-571d1843-96d1-497b-a53c-2cc03d815194",
                "notifications_enabled": true
            },
            {
                "can_call": true,
                "can_send_message": true,
                "chat_type": "direct",
                "counters_enabled": true,
                "created": "2019-09-18T00:00:25.083140Z",
                "display_name": "Контакт №21",
                "gentime": 1568764825085393152,
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/f08c50/21/256.png"
                },
                "jid": "d-14ba8969-e3a9-4daf-b6c3-bed4ea931e06",
                "notifications_enabled": true
            },
            {
                "can_call": true,
                "can_send_message": true,
                "chat_type": "direct",
                "counters_enabled": true,
                "created": "2019-09-18T00:00:23.970206Z",
                "display_name": "Контакт №20",
                "gentime": 1568764823973247744,
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/f0769a/20/256.png"
                },
                "jid": "d-13b47a4f-4473-4a8f-82ef-00b00ce5c02f",
                "notifications_enabled": true
            },
            {
                "can_call": true,
                "can_send_message": true,
                "chat_type": "direct",
                "counters_enabled": true,
                "created": "2019-09-18T00:00:23.013194Z",
                "display_name": "Контакт №19",
                "gentime": 1568764823017283072,
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/e36659/19/256.png"
                },
                "jid": "d-1d60b554-5da7-4f17-a078-eb77168081a8",
                "notifications_enabled": true
            },
            {
                "can_call": true,
                "can_send_message": true,
                "chat_type": "direct",
                "counters_enabled": true,
                "created": "2019-09-18T00:00:21.974534Z",
                "display_name": "Контакт №18",
                "gentime": 1568764821976428800,
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/8e88eb/18/256.png"
                },
                "jid": "d-ce8117ae-867d-400d-b7c8-e7df6a2ecee7",
                "notifications_enabled": true
            },
            {
                "can_call": true,
                "can_send_message": true,
                "chat_type": "direct",
                "counters_enabled": true,
                "created": "2019-09-18T00:00:21.113294Z",
                "display_name": "Контакт №17",
                "gentime": 1568764821115754496,
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/579dda/17/256.png"
                },
                "jid": "d-951775af-15fd-4e5f-bb13-099c9afd6586",
                "notifications_enabled": true
            },
            {
                "can_call": true,
                "can_send_message": true,
                "chat_type": "direct",
                "counters_enabled": true,
                "created": "2019-09-18T00:00:20.254806Z",
                "display_name": "Контакт №16",
                "gentime": 1568764820257041664,
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/63bed3/16/256.png"
                },
                "jid": "d-c540dd6e-4da3-4b65-8e89-bd35554978c2",
                "notifications_enabled": true
            },
            {
                "can_call": true,
                "can_send_message": true,
                "chat_type": "direct",
                "counters_enabled": true,
                "created": "2019-09-18T00:00:19.406973Z",
                "display_name": "Контакт №15",
                "gentime": 1568764819409007360,
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/578e3c/15/256.png"
                },
                "jid": "d-6496b27c-3251-4683-896a-a7f2191b7a33",
                "notifications_enabled": true
            },
            {
                "can_call": true,
                "can_send_message": true,
                "chat_type": "direct",
                "counters_enabled": true,
                "created": "2019-09-18T00:00:18.430679Z",
                "display_name": "Контакт №14",
                "gentime": 1568764818432459008,
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/f08c50/14/256.png"
                },
                "jid": "d-84e3cb48-598b-4ffd-918e-0662ebeb3469",
                "notifications_enabled": true
            },
            {
                "can_call": true,
                "can_send_message": true,
                "chat_type": "direct",
                "counters_enabled": true,
                "created": "2019-09-18T00:00:17.577868Z",
                "display_name": "Контакт №13",
                "gentime": 1568764817580576512,
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/f0769a/13/256.png"
                },
                "jid": "d-26879ae0-a4b9-4523-8875-4e26c4787d27",
                "notifications_enabled": true
            },
            {
                "can_call": true,
                "can_send_message": true,
                "chat_type": "direct",
                "counters_enabled": true,
                "created": "2019-09-18T00:00:16.423089Z",
                "display_name": "Контакт №12",
                "gentime": 1568764816433606656,
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/e36659/12/256.png"
                },
                "jid": "d-5049173a-93dc-4904-8796-6807b6f07bf1",
                "notifications_enabled": true
            },
            {
                "can_call": true,
                "can_send_message": true,
                "chat_type": "direct",
                "counters_enabled": true,
                "created": "2019-09-18T00:00:15.370702Z",
                "display_name": "Контакт №11",
                "gentime": 1568764815376827136,
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/8e88eb/11/256.png"
                },
                "jid": "d-8bb2ed5e-5cf4-457c-9a7b-070d7e4116a9",
                "notifications_enabled": true
            },
            {
                "can_call": true,
                "can_send_message": true,
                "chat_type": "direct",
                "counters_enabled": true,
                "created": "2019-09-18T00:00:14.302001Z",
                "display_name": "Контакт №10",
                "gentime": 1568764814322221056,
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/579dda/10/256.png"
                },
                "jid": "d-a87e168a-78df-49aa-9abb-9b5409eb7de5",
                "notifications_enabled": true
            },
            {
                "can_call": true,
                "can_send_message": true,
                "chat_type": "direct",
                "counters_enabled": true,
                "created": "2019-09-18T00:00:13.306725Z",
                "display_name": "Контакт №9",
                "gentime": 1568764813310813440,
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/63bed3/%D0%9A%E2%84%96/256.png"
                },
                "jid": "d-425ef763-2162-4dcf-8e0e-5ce0615518d8",
                "notifications_enabled": true
            },
            {
                "can_call": true,
                "can_send_message": true,
                "chat_type": "direct",
                "counters_enabled": true,
                "created": "2019-09-18T00:00:12.378013Z",
                "display_name": "Контакт №8",
                "gentime": 1568764812385766400,
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/578e3c/%D0%9A%E2%84%96/256.png"
                },
                "jid": "d-2bde58fa-3b2b-4c10-8041-b0a65e6d74b4",
                "notifications_enabled": true
            },
            {
                "can_call": true,
                "can_send_message": true,
                "chat_type": "direct",
                "counters_enabled": true,
                "created": "2019-09-18T00:00:11.451771Z",
                "display_name": "Контакт №7",
                "gentime": 1568764811455242240,
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/f08c50/%D0%9A%E2%84%96/256.png"
                },
                "jid": "d-f73f2847-4fcb-4f7e-bc11-8d2aaf21cf00",
                "notifications_enabled": true
            },
            {
                "can_call": true,
                "can_send_message": true,
                "chat_type": "direct",
                "counters_enabled": true,
                "created": "2019-09-18T00:00:10.605227Z",
                "display_name": "Контакт №6",
                "gentime": 1568764810616535296,
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/f0769a/%D0%9A%E2%84%96/256.png"
                },
                "jid": "d-af8d0ae3-ff07-4c49-94d0-7d531f972b0e",
                "notifications_enabled": true
            },
            {
                "can_call": true,
                "can_send_message": true,
                "chat_type": "direct",
                "counters_enabled": true,
                "created": "2019-09-18T00:00:09.803211Z",
                "display_name": "Контакт №5",
                "gentime": 1568764809805365248,
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/e36659/%D0%9A%E2%84%96/256.png"
                },
                "jid": "d-53e898d1-3d7b-4cc4-9ff6-ca034b5709cb",
                "notifications_enabled": true
            },
            {
                "can_call": true,
                "can_send_message": true,
                "chat_type": "direct",
                "counters_enabled": true,
                "created": "2019-09-18T00:00:09.082753Z",
                "display_name": "Контакт №4",
                "gentime": 1568764809084504320,
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/8e88eb/%D0%9A%E2%84%96/256.png"
                },
                "jid": "d-c094f142-1b54-4fae-8128-9b41504b56a9",
                "notifications_enabled": true
            },
            {
                "can_call": true,
                "can_send_message": true,
                "chat_type": "direct",
                "counters_enabled": true,
                "created": "2019-09-18T00:00:08.302823Z",
                "display_name": "Контакт №3",
                "gentime": 1568764808307773184,
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/579dda/%D0%9A%E2%84%96/256.png"
                },
                "jid": "d-ee3d7c07-fe0f-4cbd-bf3a-b8e72f0ff1c6",
                "notifications_enabled": true
            },
            {
                "can_send_message": true,
                "chat_type": "direct",
                "counters_enabled": true,
                "created": "2019-09-18T00:00:07.568804Z",
                "display_name": "tada.teamBot",
                "gentime": 1568764899109919304,
                "icons": {
                    "lg": {
                        "height": 512,
                        "url": "http://127.0.0.1:8000/static/tada-bots/systembot512.png",
                        "width": 512
                    },
                    "sm": {
                        "height": 256,
                        "url": "http://127.0.0.1:8000/static/tada-bots/systembot256.png",
                        "width": 256
                    }
                },
                "jid": "d-23eae5d4-4371-4a43-8815-a6a655cb623e",
                "last_message": {
                    "chat": "d-23eae5d4-4371-4a43-8815-a6a655cb623e",
                    "chat_type": "direct",
                    "content": {
                        "text": "Новый участник команды: @Контакт_№32",
                        "type": "plain"
                    },
                    "created": "2019-09-18T00:01:39.103766Z",
                    "from": "d-23eae5d4-4371-4a43-8815-a6a655cb623e",
                    "gentime": 1568764899109919304,
                    "is_last": true,
                    "links": [
                        {
                            "_notify_ids": [
                                34
                            ],
                            "_status": 1,
                            "pattern": "@Контакт_№32",
                            "text": "@Контакт №32",
                            "url": "otv://d-a792832a-d9f6-4bbc-b545-f2800e329c8c"
                        }
                    ],
                    "message_id": "6610b580-1d97-46da-902c-36fe03db9392",
                    "prev": "27980080-6d7c-4389-b2df-958bac09c801",
                    "to": "d-bd500a50-3a38-44d1-a43c-fb1a48e1a79e"
                },
                "notifications_enabled": true,
                "num_unread": 31
            },
            {
                "can_send_message": true,
                "chat_type": "direct",
                "counters_enabled": true,
                "created": "2019-09-18T00:00:07.517104Z",
                "display_name": "Мои заметки",
                "gentime": 1568764807520538368,
                "icons": {
                    "lg": {
                        "height": 512,
                        "url": "http://127.0.0.1:8000/static/tada-bots/notebot512.png",
                        "width": 512
                    },
                    "sm": {
                        "height": 256,
                        "url": "http://127.0.0.1:8000/static/tada-bots/notebot256.png",
                        "width": 256
                    }
                },
                "jid": "d-6d02dd61-8c07-4d11-a1f5-4584569d7b6d",
                "notifications_enabled": true
            },
            {
                "can_call": true,
                "can_send_message": true,
                "chat_type": "direct",
                "counters_enabled": true,
                "created": "2019-09-18T00:00:07.435409Z",
                "display_name": "Контакт №1",
                "gentime": 1568764807438561792,
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/578e3c/%D0%9A%E2%84%96/256.png"
                },
                "jid": "d-181b9859-ea3f-4e26-85c9-9e7ee7aaaa9a",
                "notifications_enabled": true
            }
        ],
        "group_sections": [],
        "groups": [
            {
                "can_call": true,
                "can_send_message": true,
                "chat_type": "group",
                "counters_enabled": true,
                "created": "2019-09-18T00:00:06.105362Z",
                "default_for_all": true,
                "description": "Сюда автоматически попадают все участники команды",
                "display_name": "Общий чат",
                "gentime": 1568764808134709248,
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/f0769a/%D0%9E%D1%87/256.png"
                },
                "jid": "g-82d6c0d4-8921-4519-abb3-02f17373112f",
                "notifications_enabled": true,
                "num_members": 32,
                "public": true,
                "status": "member"
            }
        ],
        "task_chats": [
            {
                "assignee": "d-bd500a50-3a38-44d1-a43c-fb1a48e1a79e",
                "can_call": true,
                "can_send_message": true,
                "chat_type": "task",
                "counters_enabled": true,
                "created": "2019-09-18T00:23:46.506612Z",
                "display_name": "1. Загрузить фото профиля",
                "gentime": 1568766227276560299,
                "icons": {
                    "lg": {
                        "height": 512,
                        "url": "http://127.0.0.1:8000/static/tada-bots/systembot512.png",
                        "width": 512
                    },
                    "sm": {
                        "height": 256,
                        "url": "http://127.0.0.1:8000/static/tada-bots/systembot256.png",
                        "width": 256
                    }
                },
                "jid": "t-2185340c-2956-4433-a4be-c42c73a5f226",
                "last_message": {
                    "chat": "t-2185340c-2956-4433-a4be-c42c73a5f226",
                    "chat_type": "task",
                    "content": {
                        "actor": "d-23eae5d4-4371-4a43-8815-a6a655cb623e",
                        "text": "Created task for @Контакт_№2: Загрузить фото профиля\nВ мобильном приложении настройки вашего профиля находятся слева, в боковом меню. В веб-версии — в правом верхнем углу.",
                        "title": "Created task for @Контакт_№2: Загрузить фото профиля\nВ мобильном приложении настройки вашего профиля находятся слева, в боковом меню. В веб-версии — в правом верхнем углу.",
                        "type": "change"
                    },
                    "created": "2019-09-18T00:23:47.276536Z",
                    "from": "d-23eae5d4-4371-4a43-8815-a6a655cb623e",
                    "gentime": 1568766227276560299,
                    "is_first": true,
                    "is_last": true,
                    "links": [
                        {
                            "_notify_ids": [
                                4
                            ],
                            "_status": 1,
                            "pattern": "@Контакт_№2",
                            "text": "@Контакт №2",
                            "url": "otv://d-bd500a50-3a38-44d1-a43c-fb1a48e1a79e"
                        }
                    ],
                    "message_id": "977f228b-2aa5-4b8b-a4f2-0ff36116aacc",
                    "notice": true,
                    "to": "t-2185340c-2956-4433-a4be-c42c73a5f226"
                },
                "notifications_enabled": true,
                "num": 1,
                "num_members": 2,
                "num_unread": 1,
                "num_unread_notices": 1,
                "observers": [],
                "owner": "d-23eae5d4-4371-4a43-8815-a6a655cb623e",
                "tabs": [
                    "active",
                    "in"
                ],
                "task_status": "new",
                "title": "Загрузить фото профиля"
            }
        ],
        "task_sections": [],
        "task_tabs": [
            {
                "filters": [
                    {
                        "field": "assignee",
                        "title": "Исполнитель"
                    },
                    {
                        "field": "owner",
                        "title": "Постановщик"
                    },
                    {
                        "field": "section",
                        "title": "Проект"
                    },
                    {
                        "field": "tag",
                        "title": "Тэг"
                    }
                ],
                "hide_empty": false,
                "key": "active",
                "pagination": false,
                "show_counter": true,
                "sort": [
                    {
                        "key": "activity",
                        "title": "Активность"
                    },
                    {
                        "key": "deadline",
                        "title": "Крайний срок"
                    },
                    {
                        "key": "new",
                        "title": "Новые"
                    },
                    {
                        "key": "old",
                        "title": "Старые"
                    }
                ],
                "title": "Все"
            },
            {
                "filters": [
                    {
                        "field": "assignee",
                        "title": "Исполнитель"
                    },
                    {
                        "field": "owner",
                        "title": "Постановщик"
                    },
                    {
                        "field": "section",
                        "title": "Проект"
                    },
                    {
                        "field": "tag",
                        "title": "Тэг"
                    }
                ],
                "hide_empty": true,
                "key": "expired",
                "pagination": false,
                "show_counter": true,
                "sort": [
                    {
                        "key": "deadline",
                        "title": "Крайний срок"
                    },
                    {
                        "key": "new",
                        "title": "Новые"
                    },
                    {
                        "key": "old",
                        "title": "Старые"
                    }
                ],
                "title": "Истёк срок"
            },
            {
                "filters": [
                    {
                        "field": "owner",
                        "title": "Постановщик"
                    },
                    {
                        "field": "section",
                        "title": "Проект"
                    },
                    {
                        "field": "tag",
                        "title": "Тэг"
                    }
                ],
                "hide_empty": false,
                "key": "in",
                "pagination": false,
                "show_counter": true,
                "sort": [
                    {
                        "key": "activity",
                        "title": "Активность"
                    },
                    {
                        "key": "deadline",
                        "title": "Крайний срок"
                    },
                    {
                        "key": "new",
                        "title": "Новые"
                    },
                    {
                        "key": "old",
                        "title": "Старые"
                    }
                ],
                "title": "Входящие"
            },
            {
                "filters": [
                    {
                        "field": "assignee",
                        "title": "Исполнитель"
                    },
                    {
                        "field": "section",
                        "title": "Проект"
                    },
                    {
                        "field": "tag",
                        "title": "Тэг"
                    }
                ],
                "hide_empty": false,
                "key": "out",
                "pagination": false,
                "show_counter": true,
                "sort": [
                    {
                        "key": "activity",
                        "title": "Активность"
                    },
                    {
                        "key": "deadline",
                        "title": "Крайний срок"
                    },
                    {
                        "key": "new",
                        "title": "Новые"
                    },
                    {
                        "key": "old",
                        "title": "Старые"
                    }
                ],
                "title": "Исходящие"
            },
            {
                "filters": [
                    {
                        "field": "assignee",
                        "title": "Исполнитель"
                    },
                    {
                        "field": "owner",
                        "title": "Постановщик"
                    },
                    {
                        "field": "section",
                        "title": "Проект"
                    },
                    {
                        "field": "tag",
                        "title": "Тэг"
                    }
                ],
                "hide_empty": true,
                "key": "copy",
                "pagination": false,
                "show_counter": true,
                "sort": [
                    {
                        "key": "activity",
                        "title": "Активность"
                    },
                    {
                        "key": "deadline",
                        "title": "Крайний срок"
                    },
                    {
                        "key": "new",
                        "title": "Новые"
                    },
                    {
                        "key": "old",
                        "title": "Старые"
                    }
                ],
                "title": "Наблюдаю"
            },
            {
                "filters": [
                    {
                        "field": "assignee",
                        "title": "Исполнитель"
                    },
                    {
                        "field": "owner",
                        "title": "Постановщик"
                    },
                    {
                        "field": "section",
                        "title": "Проект"
                    },
                    {
                        "field": "tag",
                        "title": "Тэг"
                    }
                ],
                "hide_empty": true,
                "key": "done",
                "pagination": true,
                "show_counter": false,
                "sort": [
                    {
                        "key": "done",
                        "title": "Дата завершения"
                    },
                    {
                        "key": "new",
                        "title": "Новые"
                    },
                    {
                        "key": "old",
                        "title": "Старые"
                    }
                ],
                "title": "Готово"
            }
        ]
    }
}`,
	"server.time": `{
    "event": "server.time",
    "params": {
        "time": "2019-09-27T01:01:33.216665Z"
    }
}`,
	"server.login": `{
    "event": "server.login",
    "params": {
        "device_name": "(unknown device)"
    }
}`,
	"server.section.updated": `{
    "event": "server.section.updated",
    "params": {
        "chat_type": "group",
        "gentime": 1569546093241191168,
        "sections": []
    }
}`,
	"server.section.deleted": `{
    "event": "server.section.deleted",
    "params": {
        "chat_type": "direct",
        "gentime": 1568786364815431424,
        "sections": [
            {
                "uid": "8daeb4a2-4863-4953-bcf1-c58b84dc5272"
            }
        ]
    }
}`,
	"server.online": `{
    "event": "server.online",
    "params": {
        "calls": [
            {
                "jid": "g-3f7e2a85-49f5-4586-8229-d9c52813dcb2"
            }
        ],
        "contacts": [
            {
                "afk": false,
                "jid": "d-bd500a50-3a38-44d1-a43c-fb1a48e1a79e",
                "mobile": false
            }
        ]
    }
}`,
	"server.contact.updated": `{
    "event": "server.contact.updated",
    "params": {
        "contacts": [
            {
                "can_add_to_group": true,
                "can_call": true,
                "can_create_task": true,
                "can_send_message": true,
                "changeable_fields": [],
                "contact_email": "",
                "contact_phone": "+75550000031",
                "display_name": "Контакт №32",
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/8e88eb/32/256.png"
                },
                "jid": "d-a792832a-d9f6-4bbc-b545-f2800e329c8c",
                "last_activity": null,
                "role": "",
                "sections": [],
                "status": "member"
            }
        ]
    }
}`,
	"server.remind.updated": `{
    "event": "server.remind.updated",
    "params": {
        "reminds": [
            {
                "chat": "t-2c527748-2e00-440e-a625-58121d42fdad",
                "fire_at": "2019-09-27T11:01:03.428564Z",
                "uid": "22018199-c3ae-4a9c-829e-985e975eb62a"
            }
        ]
    }
}`,
	"server.remind.deleted": `{
    "event": "server.remind.deleted",
    "params": {
        "reminds": [
            {
                "uid": "22018199-c3ae-4a9c-829e-985e975eb62a"
            }
        ]
    }
}`,
	"server.remind.fired": `{
    "event": "server.remind.fired",
    "params": {
        "reminds": [
            {
                "chat": "t-2c527748-2e00-440e-a625-58121d42fdad",
                "fire_at": "2019-09-27T11:01:03.428564Z",
                "uid": "22018199-c3ae-4a9c-829e-985e975eb62a"
            }
        ]
    }
}`,
	"server.chat.composing": `{
    "event": "server.chat.composing",
    "params": {
        "actor": "d-bd500a50-3a38-44d1-a43c-fb1a48e1a79e",
        "composing": true,
        "is_audio": false,
        "jid": "d-bd500a50-3a38-44d1-a43c-fb1a48e1a79e"
    }
}`,
	"server.call.talking": `{
    "event": "server.call.talking",
    "params": {
        "_level": null,
        "actor": "d-bd500a50-3a38-44d1-a43c-fb1a48e1a79e",
        "jid": "d-bd500a50-3a38-44d1-a43c-fb1a48e1a79e",
        "talking": true
    }
}`,
	"server.chat.lastread": `{
    "event": "server.chat.lastread",
    "params": {
        "badge": 0,
        "chats": [
            {
                "chat_type": "direct",
                "gentime": 1569546084300075149,
                "jid": "d-2a9d8c43-41d1-479c-9c8c-f029f799a724",
                "last_read_message_id": "f4f1f09c-f978-4b54-a321-b4c82604010d",
                "num_unread": 0,
                "num_unread_notices": 0
            }
        ],
        "team_unread": {
            "direct": {
                "chats": 0,
                "messages": 0
            },
            "group": {
                "chats": 0,
                "messages": 0
            },
            "task": {
                "chats": 0,
                "messages": 0
            }
        }
    }
}`,
	"server.chat.updated": `{
    "event": "server.chat.updated",
    "params": {
        "badge": 449,
        "chats": [
            {
                "can_call": true,
                "can_send_message": true,
                "chat_type": "group",
                "counters_enabled": true,
                "created": "2019-09-18T00:00:06.105362Z",
                "default_for_all": true,
                "description": "Сюда автоматически попадают все участники команды",
                "display_name": "Общий чат",
                "gentime": 1568764808134709248,
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/f0769a/%D0%9E%D1%87/256.png"
                },
                "jid": "g-82d6c0d4-8921-4519-abb3-02f17373112f",
                "notifications_enabled": true,
                "num_members": 32,
                "public": true,
                "status": "member"
            }
        ],
        "team_unread": {
            "direct": {
                "chats": 1,
                "messages": 31
            },
            "group": {
                "chats": 0,
                "messages": 0
            },
            "task": {
                "chats": 1,
                "messages": 1
            }
        }
    }
}`,
	"server.chat.deleted": `{
    "event": "server.chat.deleted",
    "params": {
        "badge": 449,
        "chats": [
            {
                "chat_type": "group",
                "gentime": 1569546094920474368,
                "is_archive": true,
                "jid": "g-3f7e2a85-49f5-4586-8229-d9c52813dcb2"
            }
        ],
        "team_unread": {
            "direct": {
                "chats": 1,
                "messages": 31
            },
            "group": {
                "chats": 0,
                "messages": 0
            },
            "task": {
                "chats": 1,
                "messages": 1
            }
        }
    }
}`,
	"server.message.updated": `{
    "event": "server.message.updated",
    "params": {
        "badge": 51,
        "chat_counters": [
            {
                "chat_type": "task",
                "gentime": 1569546093662061959,
                "jid": "t-be962c01-14ae-4c59-aeb0-f0ff5cffab74",
                "last_read_message_id": null,
                "num_unread": 0,
                "num_unread_notices": 0
            }
        ],
        "delayed": false,
        "messages": [
            {
                "chat": "t-be962c01-14ae-4c59-aeb0-f0ff5cffab74",
                "chat_type": "task",
                "content": {
                    "actor": "d-0b1b2362-fb69-4dd6-8bfb-90c89517168d",
                    "text": "Создана задача для @Контакт_№394: Тест",
                    "title": "Создана задача для @Контакт_№394: Тест",
                    "type": "change"
                },
                "created": "2019-09-27T01:01:33.661466Z",
                "editable_until": "2019-09-28T01:01:33.661466Z",
                "from": "d-a2d94e8e-74aa-4b0b-b55d-0b4fc1ce07f4",
                "gentime": 1569546093662061959,
                "is_first": true,
                "is_last": true,
                "message_id": "e0fca4c2-0d17-4317-96cd-703c2ca781e4",
                "to": "t-be962c01-14ae-4c59-aeb0-f0ff5cffab74"
            }
        ],
        "team_unread": {
            "direct": {
                "chats": 1,
                "messages": 51
            },
            "group": {
                "chats": 0,
                "messages": 0
            },
            "task": {
                "chats": 0,
                "messages": 0
            }
        }
    }
}`,
	"server.message.received": `{
    "event": "server.message.received",
    "params": {
        "messages": [
            {
                "chat": "t-be962c01-14ae-4c59-aeb0-f0ff5cffab74",
                "message_id": "e0fca4c2-0d17-4317-96cd-703c2ca781e4",
                "num_received": 0,
                "received": false
            }
        ]
    }
}`,
	"server.call.state": `{
    "event": "server.call.state",
    "params": {
        "audiorecord": true,
        "buzz": false,
        "finish": null,
        "jid": "g-3f7e2a85-49f5-4586-8229-d9c52813dcb2",
        "onliners": [
            {
                "devices": [
                    {
                        "muted": false,
                        "useragent": "Chrome"
                    }
                ],
                "jid": "d-d6de9748-0bcd-4145-b1ce-3d7f41c1c26f",
                "muted": false
            }
        ],
        "start": "2019-09-27T01:01:35.264095Z"
    }
}`,
	"server.call.answer": `{
    "event": "server.call.answer",
    "params": {
        "candidates": [
            {
                "candidate": "candidate:2 1 tcp 1019216383 95.213.181.86 0 typ host tcptype active",
                "sdpMLineIndex": 0
            }
        ],
        "jid": "g-3f7e2a85-49f5-4586-8229-d9c52813dcb2",
        "jsep": {
            "sdp": ".....SDP....",
            "type": "answer"
        },
        "onliners": [
            {
                "devices": [
                    {
                        "browser": null,
                        "muted": true
                    },
                    {
                        "browser": null,
                        "muted": true
                    }
                ],
                "jid": "d-ef53637c-f44c-4f49-9ffb-05436eb995ce",
                "muted": true
            }
        ]
    }
}`,
	"server.call.buzz": `{
    "event": "server.call.buzz",
    "params": {
        "display_name": "имя фамилия",
        "icons": {
            "lg": {
                "height": 512,
                "url": "http://127.0.0.1:8000/static/tada-bots/systembot512.png",
                "width": 512
            },
            "sm": {
                "height": 256,
                "url": "http://127.0.0.1:8000/static/tada-bots/systembot256.png",
                "width": 256
            }
        },
        "jid": "g-3f7e2a85-49f5-4586-8229-d9c52813dcb2",
        "team": "6aefcf3b-e61c-4f35-8b5b-9d762a6a6cf9"
    }
}`,
	"server.call.buzzcancel": `{
    "event": "server.call.buzzcancel",
    "params": {
        "jid": "d-0bdfbbf5-abfa-4ed2-9f98-991d5bb80127",
        "team": "848cc926-3048-44b3-a9ba-3195a394351d"
    }
}`,
	"server.team.updated": `{
    "event": "server.team.updated",
    "params": {
        "teams": [
            {
                "can_manage_sections": true,
                "default_task_deadline": "18:00:00",
                "icons": {
                    "stub": "http://127.0.0.1:8000/u/f0769a/00/256.png"
                },
                "last_active": false,
                "max_message_update_age": 86400,
                "me": {
                    "alt_send": false,
                    "can_add_to_group": true,
                    "can_create_task": true,
                    "can_delete": true,
                    "changeable_fields": [
                        "alt_send",
                        "contact_email",
                        "contact_mshort_view",
                        "contact_phone",
                        "contact_short_view",
                        "contact_show_archived",
                        "debug_show_activity",
                        "default_lang",
                        "family_name",
                        "given_name",
                        "group_mshort_view",
                        "group_notifications_enabled",
                        "group_short_view",
                        "icons",
                        "phone",
                        "role",
                        "task_mshort_view",
                        "task_notifications_enabled",
                        "task_short_view",
                        "unread_first"
                    ],
                    "contact_email": "",
                    "contact_mshort_view": false,
                    "contact_phone": "+75550000001",
                    "contact_short_view": false,
                    "contact_show_archived": false,
                    "debug_show_activity": false,
                    "default_lang": null,
                    "display_name": "Контакт №2",
                    "family_name": "№2",
                    "given_name": "Контакт",
                    "group_mshort_view": false,
                    "group_notifications_enabled": true,
                    "group_short_view": false,
                    "icons": {
                        "stub": "http://127.0.0.1:8000/u/63bed3/%D0%9A%E2%84%96/256.png"
                    },
                    "jid": "d-bd500a50-3a38-44d1-a43c-fb1a48e1a79e",
                    "last_activity": null,
                    "role": "",
                    "sections": [],
                    "status": "member",
                    "task_mshort_view": false,
                    "task_notifications_enabled": true,
                    "task_short_view": false,
                    "unread_first": false
                },
                "name": "T-32 2019-09-18 00:00:05.603182+00:00",
                "need_confirmation": true,
                "uid": "c9d8a896-a2b6-40a1-869e-2ecc0ef2436b",
                "unread": {
                    "direct": {
                        "chats": 1,
                        "messages": 31
                    },
                    "group": {
                        "chats": 0,
                        "messages": 0
                    },
                    "task": {
                        "chats": 1,
                        "messages": 1
                    }
                }
            }
        ]
    }
}`,
	"server.team.deleted": `{
    "event": "server.team.deleted",
    "params": {
        "teams": [
            {
                "is_archive": true,
                "uid": "c9d8a896-a2b6-40a1-869e-2ecc0ef2436b"
            }
        ]
    }
}`,
	"server.team.counters": `{
    "event": "server.team.counters",
    "params": {
        "teams": [
            {
                "uid": "b3e92407-41d0-456a-9b1d-3d6b79bfa7d3",
                "unread": {
                    "direct": {
                        "chats": 1,
                        "messages": 161
                    },
                    "group": {
                        "chats": 0,
                        "messages": 0
                    },
                    "task": {
                        "chats": 1,
                        "messages": 1
                    }
                }
            },
            {
                "uid": "8b05a648-d43a-4d17-9a0e-c2a119453494",
                "unread": {
                    "direct": {
                        "chats": 0,
                        "messages": 0
                    },
                    "group": {
                        "chats": 0,
                        "messages": 0
                    },
                    "task": {
                        "chats": 1,
                        "messages": 1
                    }
                }
            },
            {
                "uid": "6aefcf3b-e61c-4f35-8b5b-9d762a6a6cf9",
                "unread": {
                    "direct": {
                        "chats": 1,
                        "messages": 3
                    },
                    "group": {
                        "chats": 0,
                        "messages": 0
                    },
                    "task": {
                        "chats": 0,
                        "messages": 0
                    }
                }
            },
            {
                "uid": "5e854825-d1ac-4b28-a389-f0e8cc3bd675",
                "unread": {
                    "direct": {
                        "chats": 1,
                        "messages": 1
                    },
                    "group": {
                        "chats": 0,
                        "messages": 0
                    },
                    "task": {
                        "chats": 0,
                        "messages": 0
                    }
                }
            },
            {
                "uid": "180b9b6d-ecb0-47aa-8712-33dbd9b4a642",
                "unread": {
                    "direct": {
                        "chats": 0,
                        "messages": 0
                    },
                    "group": {
                        "chats": 1,
                        "messages": 58
                    },
                    "task": {
                        "chats": 0,
                        "messages": 0
                    }
                }
            },
            {
                "uid": "99cabb3b-d4fd-464d-b33d-144b67e3c1df",
                "unread": {
                    "direct": {
                        "chats": 1,
                        "messages": 127
                    },
                    "group": {
                        "chats": 0,
                        "messages": 0
                    },
                    "task": {
                        "chats": 1,
                        "messages": 1
                    }
                }
            },
            {
                "uid": "523e7bdc-e73f-4150-b6b9-656d6178ba51",
                "unread": {
                    "direct": {
                        "chats": 1,
                        "messages": 63
                    },
                    "group": {
                        "chats": 0,
                        "messages": 0
                    },
                    "task": {
                        "chats": 1,
                        "messages": 1
                    }
                }
            },
            {
                "uid": "c9d8a896-a2b6-40a1-869e-2ecc0ef2436b",
                "unread": {
                    "direct": {
                        "chats": 1,
                        "messages": 31
                    },
                    "group": {
                        "chats": 0,
                        "messages": 0
                    },
                    "task": {
                        "chats": 1,
                        "messages": 1
                    }
                }
            }
        ]
    }
}`,
	"client.ping": `{
	"confirm_id": "8aad294579b8",
	"event": "client.ping"
}`,
	"client.activity": `{
	"confirm_id": "75a406625c58",
	"event": "client.activity",
	"params": {
		"afk": "BOOL"
	}
}`,
	"client.confirm": `{
	"event": "client.confirm",
	"params": {
		"confirm_id": "str"
	}
}`,
	"client.message.updated": `{
	"confirm_id": "2694a2864526",
	"event": "client.message.updated",
	"params": {
		"comment": "STRING",
		"content": {
			"text": "...",
			"type": "plain"
		},
		"important": "BOOL",
		"linked_messages": [
			"STRING"
		],
		"message_id": "STRING",
		"nopreview": "BOOL",
		"to": "JID"
	}
}`,
	"client.message.deleted": `{
	"confirm_id": "cd778785149a",
	"event": "client.message.deleted",
	"params": {
		"message_id": "STRING"
	}
}`,
	"client.chat.composing": `{
	"confirm_id": "2bd5afaf39af",
	"event": "client.chat.composing",
	"params": {
		"jid": "JID"
	}
}`,
	"client.chat.lastread": `{
	"confirm_id": "8561d892f3d8",
	"event": "client.chat.lastread",
	"params": {
		"jid": "JID",
		"last_read_message_id": "STRING"
	}
}`,
	"client.call.sound": `{
	"confirm_id": "4a24b770a659",
	"event": "client.call.sound",
	"params": {
		"jid": "JID",
		"muted": "BOOL"
	}
}`,
	"client.call.reject": `{
	"confirm_id": "55e8cc25d534",
	"event": "client.call.reject",
	"params": {
		"jid": "JID"
	}
}`,
	"client.call.buzzcancel": `{
	"confirm_id": "8c52201ff7ed",
	"event": "client.call.buzzcancel",
	"params": {
		"jid": "JID"
	}
}`,
	"client.call.buzz": `{
	"confirm_id": "64977b08d763",
	"event": "client.call.buzz",
	"params": {
		"jid": "JID"
	}
}`,
	"client.call.offer": `{
	"confirm_id": "b45fdc034116",
	"event": "client.call.offer",
	"params": {
		"jid": "JID",
		"muted": "BOOL",
		"sdp": "STRING",
		"trickle": "BOOL"
	}
}`,
	"client.call.trickle": `{
	"confirm_id": "5bde78b37316",
	"event": "client.call.trickle",
	"params": {
		"candidate": "STRING",
		"jid": "JID",
		"sdp_mid": "STRING",
		"sdp_mline_index": "INT"
	}
}`,
	"client.call.leave": `{
	"confirm_id": "f5b6d29013c3",
	"event": "client.call.leave",
	"params": {
		"jid": "JID",
		"reason": "STRING"
	}
}`,
}
