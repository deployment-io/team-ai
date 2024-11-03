package agents

//type assistantWrapper struct {
//	sync.Mutex
//	assistantMap map[string]*Assistant
//}
//
//var assistantContainer = assistantWrapper{
//	assistantMap: make(map[string]*Assistant),
//}
//
//func Get(organizationID string) (*Assistant, error) {
//	assistant, ok := assistantContainer.assistantMap[organizationID]
//	if !ok {
//		assistantContainer.Lock()
//		defer assistantContainer.Unlock()
//		assistant, ok = assistantContainer.assistantMap[organizationID]
//		if !ok {
//			var err error
//			assistant, err = New()
//			if err != nil {
//				return nil, err
//			}
//			assistantContainer.assistantMap[organizationID] = assistant
//			return assistant, nil
//		}
//	}
//	return assistant, nil
//}
