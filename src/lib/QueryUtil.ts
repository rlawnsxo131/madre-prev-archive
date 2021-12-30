import { EntityTarget, Repository } from 'typeorm';

export default class QueryUtil<T> {
  private repository: Repository<T>;
  private entity: EntityTarget<T>;

  constructor(repository: Repository<T>, entity: EntityTarget<T>) {
    this.repository = repository;
    this.entity = entity;
  }

  paging() {
    return this.repository
      .createQueryBuilder()
      .from(this.entity, 'asdf')
      .where('asdf.created_at >= :');
  }
}
