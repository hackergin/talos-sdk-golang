/**
 * Copyright 2018, Xiaomi.
 * All rights reserved.
 * Author: wangfan8@xiaomi.com  
*/

package consumer


type TalosMessageReaderFactory struct {

}

func (f *TalosMessageReaderFactory) CreateMessageReader(config *TalosConsumerConfig) *TalosMessageReader {
  return NewTalosMessageReader(config)
}
