package evm

//// HandleMsgExec implements modules.AuthzMessageModule
//func (m *Module) HandleMsgExec(index int, _ *authz.MsgExec, _ int, executedMsg sdk.Msg, tx *juno.Tx) error {
//	return m.HandleMsg(index, executedMsg, tx)
//}
//
//// HandleMsg implements modules.MessageModule
//func (m *Module) HandleMsg(index int, msg sdk.Msg, tx *juno.Tx) error {
//	if len(tx.Logs) == 0 {
//		return nil
//	}
//
//	switch cosmosMsg := msg.(type) {
//	case *storagetypes.MsgCreateGroup:
//		return m.handleMsgCreateGroup(tx, cosmosMsg)
//	case *storagetypes.MsgCreateBucket:
//		return m.handleMsgCreateBucket(tx, cosmosMsg)
//	case *storagetypes.MsgCreateObject:
//		return m.handleMsgCreateObject(tx, cosmosMsg)
//	}
//
//	return nil
//}
//
//func (m *Module) handleMsgCreateGroup(tx *juno.Tx, msg *storagetypes.MsgCreateGroup) error {
//
//	// Store the group
//	groupsObj := types.NewGroups(msg.GroupId, msg.GroupName, msg.Owner, msg.SourceType, msg.Extra, msg.Tags)
//	err = m.db.SaveGroups([]types.Groups{groupsObj})
//	if err != nil {
//		return err
//	}
//}
//
//func (m *Module) handleMsgCreateBucket(tx *juno.Tx, msg *storagetypes.MsgCreateBucket) error {
//
//	bucketsObj := types.NewBuckets(msg.BucketId, msg.BucketName, msg.Owner, msg.Visibility,
//		msg.SourceType, msg.CreateAt, msg.PaymentAddress, msg.BucketStatus, msg.Tags,
//		msg.ChargedReadQuota, msg.GlobalVirtualGroupFamilyId, msg.SpAsDelegatedAgentDisabled)
//
//	return m.db.SaveBuckets([]types.Buckets{bucketsObj})
//}
//
//func (m *Module) handleMsgCreateObject(tx *juno.Tx, msg *storagetype.MsgCreateObject) error {
//
//	objects := types.NewObjects(msg.ObjectId, msg.ObjectName, msg.BucketName, msg.Owner, msg.Creator,
//		msg.PayloadSize, msg.Visibility, msg.ContentType, msg.ObjectStatus, msg.RedundancyType, msg.SourceType,
//		msg.Checksums, msg.Tags, msg.IsUpdating, msg.CreateAt, msg.UpdatedAt, msg.UpdatedBy, msg.Version, msg.LocalVirtualGroupId)
//
//	return m.db.SaveObjects([]types.Objects{objects})
//}
