import { Service } from 'typedi';
import { InjectRepository } from 'typeorm-typedi-extensions';
import { DataQueryRepository, DataRepository } from '..';
import { CreateDataParams } from '../interface/data.interface';

@Service()
export default class DataService {
  constructor(
    @InjectRepository(DataRepository)
    private readonly dataRepository: DataRepository,
    @InjectRepository(DataQueryRepository)
    private readonly dataQueryRepository: DataQueryRepository,
  ) {}

  findAll() {
    return this.dataQueryRepository.findAll();
  }

  findOne(id: string) {
    return this.dataQueryRepository.findOneById(id);
  }

  create(params: CreateDataParams) {
    return this.dataRepository.createOne(params);
  }
}
